package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(pc *PageController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1/page")
	{
		api.POST("/create", pc.CreatePage)
		api.PUT("/:id", pc.UpdatePage)
		api.GET("/:id", pc.GetPage)
		api.DELETE("/:id", pc.DeletePage)
	}

	return r
}
