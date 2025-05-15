package main

import (
	"log"
	"user-service/internal/controller"
	"user-service/internal/service"
)

func main() {
	// 初始化服务
	userService := service.NewUserService()

	// 初始化控制器
	userController := controller.NewUserController(userService)

	// 配置路由
	r := controller.SetupRouter(userController)

	// 启动服务
	log.Println("User Service starting on :8080")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
