package main

import (
	"analytics-service/internal/controller"
	"analytics-service/internal/service"
	"log"
)

func main() {
	// 初始化服务
	analyticsService := service.NewAnalyticsService()

	// 初始化控制器
	analyticsController := controller.NewAnalyticsController(analyticsService)

	// 配置路由
	r := controller.SetupRouter(analyticsController)

	// 启动服务
	log.Println("Analytics Service starting on :8087")
	err := r.Run(":8087")
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
