package main

import (
	"log"
	// 新增：导入net/http包用于HTTP服务器
	"product-service/internal/controller"
	"product-service/internal/service"
	// 新增：导入gin包
)

func main() {
	// 初始化服务
	productService := service.NewProductService()

	// 初始化控制器
	productController := controller.NewProductController(productService)

	// 配置路由
	r := controller.SetupRouter(productController)

	// 启动服务
	log.Println("Product Service starting on :8082")
	err := r.Run(":8082")
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
