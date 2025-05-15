package main

import (
	"log"
	"payment-service/internal/controller"
	"payment-service/internal/service"
)

func main() {
	// 初始化服务
	paymentService := service.NewPaymentService()

	// 初始化控制器
	paymentController := controller.NewPaymentController(paymentService)

	// 配置路由
	r := controller.SetupRouter(paymentController)

	// 启动服务
	log.Println("Payment Service starting on :8084")
	err := r.Run(":8084")
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
