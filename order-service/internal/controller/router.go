package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(oc *OrderController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1/order")
	{
		api.POST("/create", oc.CreateOrder)
		api.POST("/pay", oc.PayOrder)
		api.GET("/:id", oc.GetOrder)
		api.GET("/user/:user_id", oc.GetUserOrders)
	}

	return r
}
