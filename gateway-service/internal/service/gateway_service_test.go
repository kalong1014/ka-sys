package service

import (
	"gateway-service/internal/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

func TestGatewayService_RouteRequest_Auth(t *testing.T) {
	// 创建模拟认证服务
	authServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "valid_token" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}))
	defer authServer.Close()

	// 创建模拟后端服务
	backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Backend response"))
	}))
	defer backendServer.Close()

	// 初始化网关服务
	gatewayService := NewGatewayService(authServer.URL)

	// 添加测试路由
	route := domain.Route{
		ID:           "test-route",
		Path:         "/api/test",
		Method:       "GET",
		ServiceName:  "test-service",
		ServiceAddr:  backendServer.URL,
		AuthRequired: true,
	}
	gatewayService.routes[route.ID] = route
	gatewayService.rateLimiters[route.ID] = rate.NewLimiter(rate.Inf, 1)

	// 创建测试请求 - 有效令牌
	req, err := http.NewRequest("GET", "/api/test", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "valid_token")

	// 记录响应
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = req

	// 执行路由
	gatewayService.RouteRequest(c)

	// 验证结果
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Backend response")

	// 创建测试请求 - 无效令牌
	req, err = http.NewRequest("GET", "/api/test", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "invalid_token")

	// 记录响应
	recorder = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(recorder)
	c.Request = req

	// 执行路由
	gatewayService.RouteRequest(c)

	// 验证结果
	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
}
