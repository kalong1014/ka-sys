package controller

import (
	"context"
	"gateway-service/internal/domain"
	"gateway-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GatewayController struct {
	gatewayService *service.GatewayService
}

func NewGatewayController(gs *service.GatewayService) *GatewayController {
	return &GatewayController{
		gatewayService: gs,
	}
}

// 创建路由
func (c *GatewayController) CreateRoute(ctx *gin.Context) {
	var req domain.CreateRouteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	route, err := c.gatewayService.CreateRoute(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "创建路由失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, route)
}

// 创建认证配置
func (c *GatewayController) CreateAuthConfig(ctx *gin.Context) {
	var req domain.CreateAuthConfigRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	authConfig, err := c.gatewayService.CreateAuthConfig(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "创建认证配置失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, authConfig)
}

// 路由请求
func (c *GatewayController) RouteRequest(ctx *gin.Context) {
	c.gatewayService.RouteRequest(ctx)
}
