package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(dc *DomainController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1/domain")
	{
		api.POST("/create", dc.CreateDomain)
		api.PUT("/:id", dc.UpdateDomain)
		api.GET("/:id", dc.GetDomain)
		api.DELETE("/:id", dc.DeleteDomain)
	}

	return r
}
