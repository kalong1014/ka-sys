package domain

import (
	"time"
)

// 订单状态
const (
	OrderStatusPending  = "pending"  // 待支付
	OrderStatusPaid     = "paid"     // 已支付
	OrderStatusFailed   = "failed"   // 支付失败
	OrderStatusCanceled = "canceled" // 已取消
)

// 订单模型
type Order struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	ProductID   string    `json:"product_id"`
	ProductName string    `json:"product_name"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"`
	PaymentType string    `json:"payment_type"` // 支付方式：alipay, wechat, card
	PaymentTime time.Time `json:"payment_time"` // 支付时间
	CardKeyID   string    `json:"card_key_id"`  // 关联的卡密ID
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// 创建订单请求
type CreateOrderRequest struct {
	UserID      string  `json:"user_id" validate:"required"`
	MerchantID  string  `json:"merchant_id" validate:"required"`
	ProductID   string  `json:"product_id" validate:"required"`
	ProductName string  `json:"product_name" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	PaymentType string  `json:"payment_type" validate:"required,oneof=alipay wechat card"`
}

// 支付订单请求
type PayOrderRequest struct {
	OrderID     string `json:"order_id" validate:"required"`
	PaymentType string `json:"payment_type" validate:"required,oneof=alipay wechat card"`
	CardKey     string `json:"card_key,omitempty"` // 卡密支付时使用
}
