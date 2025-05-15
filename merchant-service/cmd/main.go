package main

import (
	"log"
	"merchant-service/internal/controller"
	"merchant-service/internal/repository"
	"merchant-service/internal/service"
)

func main() {
	// 从环境变量获取数据库连接信息
	dsn := "user:password@tcp(localhost:3306)/merchant_db?charset=utf8mb4&parseTime=True&loc=Local"

	// 初始化数据库仓库
	repo, err := repository.NewMerchantMySQLRepo(dsn)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 初始化服务
	merchantService := service.NewMerchantService(repo)

	// 初始化控制器
	merchantController := controller.NewMerchantController(merchantService)

	// 配置路由
	r := controller.SetupRouter(merchantController)

	// 启动服务
	log.Println("Merchant Service starting on :8081")
	err = r.Run(":8081")
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
