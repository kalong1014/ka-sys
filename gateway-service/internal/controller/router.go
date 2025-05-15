package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(gc *GatewayController) (*gin.Engine, *gin.Engine) {
	// 创建管理API的引擎（用于内部管理）
	adminEngine := gin.Default()
	adminGroup := adminEngine.Group("/api/admin")
	{
		gatewayAdmin := adminGroup.Group("/gateway")
		gatewayAdmin.POST("/routes", gc.CreateRoute)
		gatewayAdmin.POST("/auth-config", gc.CreateAuthConfig)
	}

	// 创建主网关引擎（用于代理所有外部请求）
	mainEngine := gin.Default()

	// 设置404处理器，确保所有未匹配的请求都被通配符路由捕获
	mainEngine.NoRoute(func(c *gin.Context) {
		gc.RouteRequest(c)
	})

	// 捕获所有HTTP方法的请求
	mainEngine.Any("/*path", gc.RouteRequest)

	return adminEngine, mainEngine
}
