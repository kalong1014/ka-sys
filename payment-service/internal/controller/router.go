package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(pc *PaymentController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1/payment")
	{
		api.POST("/create", pc.CreatePayment)
		api.POST("/process-result", pc.ProcessPaymentResult)
		api.POST("/refund", pc.ProcessRefund)
		api.GET("/:id", pc.GetPayment)
		api.GET("/order/:order_id", pc.GetOrderPayments)
	}

	return r
}
