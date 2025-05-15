package main

import (
	"log"
	"merchant-service/internal/controller"
	"merchant-service/internal/service"
)

func main() {
	// 初始化服务
	merchantService := service.NewMerchantService()

	// 初始化控制器
	merchantController := controller.NewMerchantController(merchantService)

	// 配置路由
	r := controller.SetupRouter(merchantController)

	// 启动服务
	log.Println("Merchant Service starting on :8081")
	err := r.Run(":8081")
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
