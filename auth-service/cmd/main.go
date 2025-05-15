package main

import (
	"auth-service/internal/controller"
	"auth-service/internal/service"
	"log"
)

const (
	jwtSecretKey  = "your_secret_key_here_12345" // 生产环境应从配置获取
	tokenLifetime = 24                           // 令牌有效期（小时）
)

func main() {
	// 初始化服务
	authService := service.NewAuthService(jwtSecretKey, tokenLifetime)

	// 初始化控制器
	authController := controller.NewAuthController(authService)

	// 配置路由
	r := controller.SetupRouter(authController)

	// 启动服务
	log.Println("Auth Service starting on :8090")
	err := r.Run(":8090")
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
