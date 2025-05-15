package main

import (
	"gateway-service/internal/controller"
	"gateway-service/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	// 从环境变量获取认证服务地址
	authServiceAddr := "http://localhost:8090"

	// 初始化服务 (删除重复声明)
	gatewayService := service.NewGatewayService(authServiceAddr)

	// 设置为生产模式
	gin.SetMode(gin.ReleaseMode)

	// 初始化控制器
	gatewayController := controller.NewGatewayController(gatewayService)

	// 配置路由
	adminEngine, mainEngine := controller.SetupRouter(gatewayController)

	// 启动管理API服务
	go func() {
		log.Println("Admin API Service starting on :8088")
		if err := adminEngine.Run(":8088"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Admin API server failed: %v", err)
		}
	}()

	// 启动主网关服务
	go func() {
		log.Println("Gateway Service starting on :8089")
		if err := mainEngine.Run(":8089"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Gateway server failed: %v", err)
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down servers...")
}
