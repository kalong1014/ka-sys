package main

import (
	"context"
	"log"
	"order-service/internal/controller"
	"order-service/internal/repository"
	"order-service/internal/service"
)

func main() {
	// 初始化RabbitMQ客户端
	mqClient, err := mq.NewRabbitMQClient(
		"amqp://guest:guest@localhost:5672/",
		"order_service",
	)
	if err != nil {
		log.Fatalf("初始化RabbitMQ失败: %v", err)
	}

	// 声明交换机和队列（示例）
	if err := mqClient.DeclareExchange("order_exchange", "direct"); err != nil {
		log.Fatalf("声明交换机失败: %v", err)
	}
	if err := mqClient.BindQueue("order_created_queue", "order_exchange", "order_created"); err != nil {
		log.Fatalf("绑定队列失败: %v", err)
	}

	// 启动消费者
	go func() {
		if err := mqClient.Consume("order_created_queue", func(message []byte) error {
			return s.HandleOrderCreated(context.Background(), message)
		}); err != nil {
			log.Fatalf("启动消费者失败: %v", err)
		}
	}()

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
