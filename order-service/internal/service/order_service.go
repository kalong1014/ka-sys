package service

import (
	"context"
	"errors"
	"log"
	"order-service/internal/domain"
	"time"

	"common/pkg/mq"
	"encoding/json"

	"github.com/google/uuid"
)

type OrderService struct {
	orders   map[string]domain.Order // 订单存储
	repo     repository.OrderRepository
	mqClient *mq.RabbitMQClient // 新增消息队列客户端
	// 后续需要注入商品服务和卡密服务的客户端
}

func NewOrderService(repo repository.OrderRepository, mqClient *mq.RabbitMQClient) *OrderService {
	return &OrderService{
		repo:     repo,
		mqClient: mqClient,
		orders:   make(map[string]domain.Order),
	}
}

// 创建订单
func (s *OrderService) CreateOrder(ctx context.Context, req *domain.CreateOrderRequest) (*domain.Order, error) {
	log.Printf("开始创建订单: 用户ID=%s, 商户ID=%s, 商品ID=%s", req.UserID, req.MerchantID, req.ProductID)

	// 生成订单ID
	orderID := uuid.New().String()

	// 创建订单对象
	order := domain.Order{
		ID:          orderID,
		UserID:      req.UserID,
		MerchantID:  req.MerchantID,
		ProductID:   req.ProductID,
		ProductName: req.ProductName,
		Price:       req.Price,
		Status:      domain.OrderStatusPending,
		PaymentType: req.PaymentType,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// 保存订单
	s.orders[orderID] = order
	log.Printf("订单创建成功: ID=%s", orderID)

	// 发送订单创建消息
	orderCreatedEvent := domain.OrderCreatedEvent{
		OrderID:     order.ID,
		MerchantID:  order.MerchantID,
		TotalAmount: order.TotalAmount,
		CreatedAt:   time.Now(),
	}

	eventData, err := json.Marshal(orderCreatedEvent)
	if err != nil {
		log.Printf("序列化订单创建事件失败: %v", err)
		// 可以选择记录错误但不影响主流程
	} else {
		err = s.mqClient.Publish(eventData)
		if err != nil {
			log.Printf("发布订单创建事件失败: %v", err)
		}
	}

	return &order, nil
}

// 支付订单
func (s *OrderService) PayOrder(ctx context.Context, req *domain.PayOrderRequest) (*domain.Order, error) {
	log.Printf("开始处理订单支付: 订单ID=%s, 支付方式=%s", req.OrderID, req.PaymentType)

	// 检查订单是否存在
	order, exists := s.orders[req.OrderID]
	if !exists {
		log.Println("订单不存在")
		return nil, errors.New("订单不存在")
	}

	// 检查订单状态
	if order.Status != domain.OrderStatusPending {
		log.Printf("订单状态不正确: 当前状态=%s", order.Status)
		return nil, errors.New("订单状态不正确")
	}

	// 检查支付方式是否匹配
	if order.PaymentType != req.PaymentType {
		log.Printf("支付方式不匹配: 订单支付方式=%s, 请求支付方式=%s", order.PaymentType, req.PaymentType)
		return nil, errors.New("支付方式不匹配")
	}

	// 处理支付逻辑（这里简化处理，实际应调用支付服务）
	// TODO: 调用支付服务完成支付

	// 更新订单状态
	order.Status = domain.OrderStatusPaid
	order.PaymentTime = time.Now()
	order.UpdatedAt = time.Now()

	// 如果是卡密支付，记录卡密ID
	if req.PaymentType == "card" {
		order.CardKeyID = req.CardKey
		// TODO: 验证卡密有效性并标记为已使用
	}

	s.orders[req.OrderID] = order
	log.Printf("订单支付成功: ID=%s", req.OrderID)

	return &order, nil
}

// 获取订单信息
func (s *OrderService) GetOrder(ctx context.Context, orderID string) (*domain.Order, error) {
	log.Printf("获取订单信息: ID=%s", orderID)

	order, exists := s.orders[orderID]
	if !exists {
		log.Println("订单不存在")
		return nil, errors.New("订单不存在")
	}

	return &order, nil
}

// 获取用户订单列表
func (s *OrderService) GetUserOrders(ctx context.Context, userID string) (*[]domain.Order, error) {
	log.Printf("获取用户订单列表: 用户ID=%s", userID)

	var userOrders []domain.Order
	for _, order := range s.orders {
		if order.UserID == userID {
			userOrders = append(userOrders, order)
		}
	}

	log.Printf("找到 %d 个订单", len(userOrders))
	return &userOrders, nil
}
