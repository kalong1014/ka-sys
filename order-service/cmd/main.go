package main

import (
	"log"
	"order-service/internal/controller"
	"order-service/internal/repository"
	"order-service/internal/service"
)

func main() {
	dsn := "user:password@tcp(localhost:3306)/order_db?charset=utf8mb4&parseTime=True&loc=Local"

	// 初始化数据库仓库
	repo, err := repository.NewOrderMySQLRepo(dsn)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 初始化服务
	orderService := service.NewOrderService(repo)

	// 初始化服务
	orderService := service.NewOrderService()

	// 初始化控制器
	orderController := controller.NewOrderController(orderService)

	// 配置路由
	r := controller.SetupRouter(orderController)

	// 启动服务
	log.Println("Order Service starting on :8083")
	err := r.Run(":8083")
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
