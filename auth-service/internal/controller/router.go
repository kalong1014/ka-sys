package controller

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter 配置路由
func SetupRouter(ac *AuthController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1/auth")
	{
		api.POST("/register", ac.Register)
		api.POST("/login", ac.Login)
		api.POST("/validate-token", ac.ValidateToken)
		api.GET("/check-permission", ac.CheckPermission)
	}

	return r
}
