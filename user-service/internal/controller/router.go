package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(uc *UserController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1/user")
	{
		api.POST("/register", uc.Register)
		api.POST("/login", uc.Login)
	}

	return r
}
