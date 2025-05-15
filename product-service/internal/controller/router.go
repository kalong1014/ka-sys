package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(pc *ProductController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1/product")
	{
		api.POST("/create", pc.CreateProduct)
		api.POST("/generate-card-keys", pc.GenerateCardKeys)
		api.GET("/card-key/:id", pc.GetCardKey)
		api.PUT("/card-key/:id/status", pc.UpdateCardKeyStatus)
	}

	return r
}
