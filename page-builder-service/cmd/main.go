package main

import (
	"log"
	"page-builder-service/internal/controller"
	"page-builder-service/internal/service"
)

func main() {
	// 初始化服务
	pageService := service.NewPageService()

	// 初始化控制器
	pageController := controller.NewPageController(pageService)

	// 配置路由
	r := controller.SetupRouter(pageController)

	// 启动服务
	log.Println("Page Builder Service starting on :8085")
	err := r.Run(":8085")
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
