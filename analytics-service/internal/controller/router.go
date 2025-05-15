package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(ac *AnalyticsController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1/analytics")
	{
		api.POST("/visit", ac.CreateVisit)
		api.POST("/order-stats", ac.GetOrderStats)
		api.POST("/traffic-stats", ac.GetTrafficStats)
	}

	return r
}
