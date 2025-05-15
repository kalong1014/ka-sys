package main

import (
	"domain-service/internal/controller"
	"domain-service/internal/service"
	"log"
)

func main() {
	// 初始化服务
	domainService := service.NewDomainService()

	// 初始化控制器
	domainController := controller.NewDomainController(domainService)

	// 配置路由
	r := controller.SetupRouter(domainController)

	// 启动服务
	log.Println("Domain Service starting on :8086")
	err := r.Run(":8086")
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
