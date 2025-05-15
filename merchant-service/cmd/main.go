package main

import (
	"fmt"
	"merchant-service/internal/controller"
	"merchant-service/internal/repository"
	"merchant-service/internal/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	config := config.LoadConfig()

	// 初始化日志
	logger := logger.InitLogger()
	ctxLogger := logger.NewContextLogger(logger)

	// 初始化服务发现客户端
	consulClient, err := discovery.NewConsulClient(config.ConsulAddress)
	if err != nil {
		ctxLogger.Error("连接Consul失败", err)
		os.Exit(1)
	}

	// 生成服务ID
	serviceID := fmt.Sprintf("%s-%s-%d", config.ServiceName, config.Host, config.Port)

	// 注册服务
	err = consulClient.RegisterService(
		serviceID,
		config.ServiceName,
		config.Host,
		config.Port,
		[]string{"merchant", "api"},
	)

	if err != nil {
		ctxLogger.Error("注册服务失败", err)
		os.Exit(1)
	}

	// 确保服务退出时注销
	defer func() {
		if err := consulClient.DeregisterService(serviceID); err != nil {
			ctxLogger.Error("注销服务失败", err)
		}
	}()

	// 初始化数据库
	db, err := database.Connect(config.DBURL)
	if err != nil {
		ctxLogger.Error("连接数据库失败", err)
		os.Exit(1)
	}

	// 初始化服务
	merchantRepo := repository.NewMerchantMySQLRepo(db)
	merchantService := service.NewMerchantService(merchantRepo, ctxLogger)

	// 初始化HTTP服务器
	router := gin.Default()

	// 添加错误处理中间件
	router.Use(errors.ErrorHandler())

	// 注册路由
	controller.NewMerchantController(merchantService).RegisterRoutes(router)

	// 启动服务器
	ctxLogger.Info("服务启动中", "port", config.Port)
	if err := router.Run(":" + config.Port); err != nil {
		ctxLogger.Error("启动服务器失败", err)
		os.Exit(1)
	}
}
