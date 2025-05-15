package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"payment-service/internal/domain"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type PaymentService struct {
	payments map[string]domain.Payment // 支付记录存储
}

func NewPaymentService() *PaymentService {
	return &PaymentService{
		payments: make(map[string]domain.Payment),
	}
}

// 创建支付
func (s *PaymentService) CreatePayment(ctx context.Context, req *domain.CreatePaymentRequest) (*domain.Payment, error) {
	log.Printf("开始创建支付: 订单ID=%s, 用户ID=%s, 金额=%.2f, 支付方式=%s",
		req.OrderID, req.UserID, req.Amount, req.Method)

	// 生成支付ID
	paymentID := uuid.New().String()

	// 创建支付对象
	payment := domain.Payment{
		ID:            paymentID,
		OrderID:       req.OrderID,
		UserID:        req.UserID,
		MerchantID:    req.MerchantID,
		Amount:        req.Amount,
		Method:        req.Method,
		Status:        domain.PaymentStatusPending,
		TransactionID: "",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// 保存支付记录
	s.payments[paymentID] = payment
	log.Printf("支付记录创建成功: ID=%s", paymentID)

	return &payment, nil
}

// 处理支付结果
func (s *PaymentService) ProcessPaymentResult(ctx context.Context, req *domain.ProcessPaymentResultRequest) (*domain.Payment, error) {
	log.Printf("开始处理支付结果: 订单ID=%s, 状态=%s", req.OrderID, req.Status)

	// 查找对应的支付记录
	var payment domain.Payment
	found := false
	for _, p := range s.payments {
		if p.OrderID == req.OrderID {
			payment = p
			found = true
			break
		}
	}

	if !found {
		log.Println("未找到对应的支付记录")
		return nil, errors.New("未找到对应的支付记录")
	}

	// 检查支付状态
	if payment.Status != domain.PaymentStatusPending {
		log.Printf("支付状态不正确: 当前状态=%s", payment.Status)
		return nil, errors.New("支付状态不正确")
	}

	// 更新支付状态
	payment.Status = req.Status
	payment.TransactionID = req.TransactionID

	// 处理支付时间
	if req.Status == domain.PaymentStatusSuccess {
		if req.PaymentTime != "" {
			t, err := time.Parse(time.RFC3339, req.PaymentTime)
			if err != nil {
				log.Printf("解析支付时间失败，使用当前时间: %v", err)
				payment.PaymentTime = time.Now()
			} else {
				payment.PaymentTime = t
			}
		} else {
			payment.PaymentTime = time.Now()
		}
	}

	payment.UpdatedAt = time.Now()
	s.payments[payment.ID] = payment

	log.Printf("支付结果处理成功: 订单ID=%s, 状态=%s", req.OrderID, req.Status)
	return &payment, nil
}

// 处理退款
func (s *PaymentService) ProcessRefund(ctx context.Context, req *domain.RefundRequest) (*domain.Payment, error) {
	log.Printf("开始处理退款: 订单ID=%s, 金额=%.2f, 原因=%s",
		req.OrderID, req.Amount, req.Reason)

	// 查找对应的支付记录
	var payment domain.Payment
	found := false
	for _, p := range s.payments {
		if p.OrderID == req.OrderID {
			payment = p
			found = true
			break
		}
	}

	if !found {
		log.Println("未找到对应的支付记录")
		return nil, errors.New("未找到对应的支付记录")
	}

	// 检查支付状态
	if payment.Status != domain.PaymentStatusSuccess {
		log.Printf("支付状态不正确，无法退款: 当前状态=%s", payment.Status)
		return nil, errors.New("支付状态不正确，无法退款")
	}

	// 检查退款金额
	if req.Amount > payment.Amount {
		log.Printf("退款金额超过支付金额: 支付金额=%.2f, 退款金额=%.2f",
			payment.Amount, req.Amount)
		return nil, errors.New("退款金额超过支付金额")
	}

	// 更新支付状态
	payment.Status = domain.PaymentStatusRefunding
	payment.RefundReason = req.Reason
	payment.UpdatedAt = time.Now()
	s.payments[payment.ID] = payment

	// 模拟退款处理（实际应调用支付渠道的退款接口）
	log.Println("调用支付渠道退款接口...")
	time.Sleep(2 * time.Second) // 模拟网络延迟

	// 更新退款状态
	payment.Status = domain.PaymentStatusRefunded
	payment.RefundTime = time.Now()
	payment.UpdatedAt = time.Now()
	s.payments[payment.ID] = payment

	log.Printf("退款处理成功: 订单ID=%s, 金额=%.2f", req.OrderID, req.Amount)
	return &payment, nil
}

func (s *PaymentService) HandleOrderCreated(ctx context.Context, eventData []byte) error {
	var event events.OrderCreatedEvent
	if err := json.Unmarshal(eventData, &event); err != nil {
		return err
	}

	// 模拟支付处理（实际应调用支付渠道API）
	paymentID, err := s.processPayment(event.OrderID, event.TotalAmount)
	if err != nil {
		// 支付失败，发送PaymentFailed事件
		failEvent := events.PaymentFailedEvent{
			OrderID: event.OrderID,
			Reason:  "模拟支付失败",
		}
		eventData, _ := json.Marshal(failEvent)
		s.mqClient.Publish("payment_failed", eventData)
		return err
	}

	// 支付成功，发送PaymentSucceeded事件
	successEvent := events.PaymentSucceededEvent{
		OrderID:   event.OrderID,
		PaymentID: paymentID,
	}
	eventData, _ := json.Marshal(successEvent)
	s.mqClient.Publish("payment_succeeded", eventData)

	return nil
}

// 模拟支付处理（示例）
func (s *PaymentService) processPayment(orderID string, amount float64) (string, error) {
	// 实际应调用支付宝/微信支付API
	if amount < 10 { // 模拟小额支付失败
		return "", errors.New("支付金额不足")
	}
	return uuid.New().String(), nil
}

// 获取支付信息
func (s *PaymentService) GetPayment(ctx context.Context, paymentID string) (*domain.Payment, error) {
	log.Printf("获取支付信息: ID=%s", paymentID)

	payment, exists := s.payments[paymentID]
	if !exists {
		log.Println("支付记录不存在")
		return nil, errors.New("支付记录不存在")
	}

	return &payment, nil
}

// 获取订单的支付记录
func (s *PaymentService) GetOrderPayments(ctx context.Context, orderID string) (*[]domain.Payment, error) {
	log.Printf("获取订单的支付记录: 订单ID=%s", orderID)

	var orderPayments []domain.Payment
	for _, payment := range s.payments {
		if payment.OrderID == orderID {
			orderPayments = append(orderPayments, payment)
		}
	}

	log.Printf("找到 %d 条支付记录", len(orderPayments))
	return &orderPayments, nil
}

// 生成交易流水号
func (s *PaymentService) generateTransactionID() string {
	// 生成规则: 时间戳 + 随机数
	timestamp := time.Now().UnixNano() / 1e6 // 毫秒级时间戳
	randomStr := uuid.New().String()[:8]     // 取UUID前8位

	return "TXN" + strconv.FormatInt(timestamp, 10) + randomStr
}
