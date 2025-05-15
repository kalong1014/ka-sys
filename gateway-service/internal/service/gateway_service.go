package service

import (
	"context"
	"errors"
	"gateway-service/internal/client" // 修改导入路径
	"gateway-service/internal/domain"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
)

type GatewayService struct {
	routes          map[string]domain.Route              // 路由配置
	authConfig      *domain.AuthConfig                   // 认证配置
	rateLimiters    map[string]*rate.Limiter             // 限流控制器
	circuitBreakers map[string]*gobreaker.CircuitBreaker // 断路器
	mu              sync.RWMutex                         // 读写锁
	authClient      *client.AuthClient                   // 修改类型
}

func NewGatewayService(authServiceAddr string) *GatewayService {
	return &GatewayService{
		routes:          make(map[string]domain.Route),
		rateLimiters:    make(map[string]*rate.Limiter),
		circuitBreakers: make(map[string]*gobreaker.CircuitBreaker),
		authClient:      client.NewAuthClient(authServiceAddr), // 修改初始化
	}
}

// 创建路由
func (s *GatewayService) CreateRoute(ctx context.Context, req *domain.CreateRouteRequest) (*domain.Route, error) {
	log.Printf("开始创建路由: 路径=%s, 服务=%s", req.Path, req.ServiceName)

	// 生成路由ID
	routeID := uuid.New().String()

	// 设置默认值
	if req.Timeout == 0 {
		req.Timeout = 5000 // 默认5秒超时
	}

	if req.RetryTimes == 0 {
		req.RetryTimes = 1 // 默认重试1次
	}

	if req.RateLimit == 0 {
		req.RateLimit = 100 // 默认100请求/秒
	}

	// 创建路由
	route := domain.Route{
		ID:           routeID,
		Path:         req.Path,
		Method:       req.Method,
		ServiceName:  req.ServiceName,
		ServiceAddr:  req.ServiceAddr,
		Timeout:      req.Timeout,
		RetryTimes:   req.RetryTimes,
		RateLimit:    req.RateLimit,
		AuthRequired: req.AuthRequired,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// 保存路由
	s.mu.Lock()
	s.routes[routeID] = route
	s.mu.Unlock()

	// 创建限流控制器
	s.rateLimiters[routeID] = rate.NewLimiter(rate.Limit(req.RateLimit), req.RateLimit)

	// 创建断路器
	s.circuitBreakers[routeID] = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        routeID,
		MaxRequests: 5,
		Interval:    60 * time.Second,
		Timeout:     10 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.ConsecutiveFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		},
	})

	log.Printf("路由创建成功: ID=%s, 路径=%s", routeID, req.Path)
	return &route, nil
}

// 创建认证配置
func (s *GatewayService) CreateAuthConfig(ctx context.Context, req *domain.CreateAuthConfigRequest) (*domain.AuthConfig, error) {
	log.Printf("开始创建认证配置: 服务=%s", req.ServiceName)

	// 生成配置ID
	configID := uuid.New().String()

	// 设置默认值
	if req.TokenHeader == "" {
		req.TokenHeader = "Authorization"
	}

	// 创建认证配置
	authConfig := domain.AuthConfig{
		ID:          configID,
		ServiceName: req.ServiceName,
		ServiceAddr: req.ServiceAddr,
		TokenHeader: req.TokenHeader,
		PublicPaths: req.PublicPaths,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// 保存认证配置
	s.mu.Lock()
	s.authConfig = &authConfig
	s.mu.Unlock()

	log.Printf("认证配置创建成功: ID=%s", configID)
	return &authConfig, nil
}

// 路由请求
func (s *GatewayService) RouteRequest(c *gin.Context) {
	path := c.Request.URL.Path
	method := c.Request.Method

	log.Printf("路由请求: 路径=%s, 方法=%s", path, method)

	// 查找匹配的路由
	route, err := s.findRoute(path, method)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到匹配的路由"})
		return
	}

	log.Printf("匹配路由: ID=%s, 目标服务=%s", route.ID, route.ServiceName)

	// 限流检查
	if !s.rateLimiters[route.ID].Allow() {
		log.Printf("请求被限流: 路由ID=%s", route.ID)
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "请求过于频繁"})
		return
	}

	// 认证检查
	if route.AuthRequired {
		s.mu.RLock()
		authConfig := s.authConfig
		s.mu.RUnlock()

		if authConfig == nil {
			log.Println("认证配置不存在，但路由需要认证")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "认证配置未初始化"})
			return
		}

		// 检查是否为公共路径
		isPublic := false
		for _, publicPath := range authConfig.PublicPaths {
			if strings.HasPrefix(path, publicPath) {
				isPublic = true
				break
			}
		}

		if !isPublic {
			// 获取令牌
			token := c.GetHeader(authConfig.TokenHeader)
			if token == "" {
				log.Println("缺少认证令牌")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
				return
			}

			// 验证令牌
			isValid, err := s.authClient.ValidateToken(token)
			if err != nil {
				log.Printf("令牌验证失败: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "令牌验证失败"})
				return
			}

			if !isValid {
				log.Println("无效令牌")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "无效令牌"})
				return
			}

			// 检查权限（从路由元数据中获取所需角色，此处简化处理）
			requiredRole := "user" // 默认需要用户权限
			hasPermission, err := s.authClient.CheckPermission(token, requiredRole)
			if err != nil {
				log.Printf("权限检查失败: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "权限检查失败"})
				return
			}

			if !hasPermission {
				log.Println("权限不足")
				c.JSON(http.StatusForbidden, gin.H{"error": "权限不足"})
				return
			}
		}
	}

	// 代理请求到目标服务
	s.proxyRequest(c, route)
}

// 查找匹配的路由
func (s *GatewayService) findRoute(path, method string) (*domain.Route, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 查找完全匹配的路由
	for _, route := range s.routes {
		if route.Path == path && route.Method == method {
			return &route, nil
		}
	}

	// 查找前缀匹配的路由（支持路径参数）
	for _, route := range s.routes {
		if strings.HasPrefix(path, route.Path) && route.Method == method {
			return &route, nil
		}
	}

	return nil, errors.New("未找到匹配的路由")
}

// 代理请求到目标服务
func (s *GatewayService) proxyRequest(c *gin.Context, route *domain.Route) {
	targetURL, err := url.Parse(route.ServiceAddr)
	if err != nil {
		log.Printf("解析目标URL失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务配置错误"})
		return
	}

	// 使用断路器包装代理请求
	_, err = s.circuitBreakers[route.ID].Execute(func() (interface{}, error) {
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		// 设置超时
		timeout := time.Duration(route.Timeout) * time.Millisecond
		proxy.Transport = &http.Transport{
			ResponseHeaderTimeout: timeout,
		}

		// 设置请求头
		c.Request.Header.Set("X-Forwarded-For", c.ClientIP())
		c.Request.Header.Set("X-Forwarded-Host", c.Request.Host)
		c.Request.Header.Set("X-Forwarded-Proto", c.Request.Proto)

		// 代理请求
		proxy.ServeHTTP(c.Writer, c.Request)
		return nil, nil
	})

	if err != nil {
		log.Printf("代理请求失败: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "服务暂时不可用"})
		return
	}
}
