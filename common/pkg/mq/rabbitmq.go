package mq

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQClient(url, queueName string) (*RabbitMQClient, error) {
	// 连接RabbitMQ（带重试机制）
	var conn *amqp.Connection
	var err error

	for i := 0; i < 5; i++ {
		conn, err = amqp.Dial(url)
		if err == nil {
			break
		}
		log.Printf("连接RabbitMQ失败 (尝试 %d/5): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	// 声明队列（持久化）
	q, err := ch.QueueDeclare(
		queueName, // 队列名称
		true,      // 是否持久化
		false,     // 是否自动删除
		false,     // 是否排他
		false,     // 是否等待服务器响应
		nil,       // 额外参数
	)

	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitMQClient{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

// 发布消息
func (c *RabbitMQClient) Publish(message []byte) error {
	return c.channel.Publish(
		"",           // 交换器
		c.queue.Name, // 路由键（队列名称）
		false,        // 是否强制
		false,        // 是否立即
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         message,
			DeliveryMode: amqp.Persistent, // 持久化消息
		},
	)
}

// 消费消息
func (c *RabbitMQClient) Consume(handler func([]byte) error) error {
	msgs, err := c.channel.Consume(
		c.queue.Name, // 队列名称
		"",           // 消费者名称
		true,         // 是否自动确认
		false,        // 是否排他
		false,        // 是否为本地队列
		false,        // 是否等待服务器响应
		nil,          // 额外参数
	)

	if err != nil {
		return err
	}

	// 启动goroutine处理消息
	go func() {
		for d := range msgs {
			if err := handler(d.Body); err != nil {
				log.Printf("处理消息失败: %v", err)
			}
		}
	}()

	return nil
}

// 关闭连接
func (c *RabbitMQClient) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
}
