package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(mc *MerchantController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1/merchant")
	{
		api.POST("/apply", mc.ApplyMerchant)
		api.POST("/review", mc.ReviewMerchant)
		api.GET("/:id", mc.GetMerchantInfo)
	}

	return r
}
