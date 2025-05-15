package domain

import (
	"time"
)

// 支付状态
const (
	PaymentStatusPending   = "pending"   // 待支付
	PaymentStatusSuccess   = "success"   // 支付成功
	PaymentStatusFailed    = "failed"    // 支付失败
	PaymentStatusRefunding = "refunding" // 退款中
	PaymentStatusRefunded  = "refunded"  // 已退款
)

// 支付方式
const (
	PaymentMethodAlipay  = "alipay"  // 支付宝
	PaymentMethodWechat  = "wechat"  // 微信支付
	PaymentMethodCard    = "card"    // 卡密支付
	PaymentMethodBalance = "balance" // 余额支付
)

// 支付记录
type Payment struct {
	ID            string    `json:"id"`
	OrderID       string    `json:"order_id"`
	UserID        string    `json:"user_id"`
	MerchantID    string    `json:"merchant_id"`
	Amount        float64   `json:"amount"`
	Method        string    `json:"method"`
	Status        string    `json:"status"`
	TransactionID string    `json:"transaction_id"` // 交易流水号
	PaymentTime   time.Time `json:"payment_time"`   // 支付时间
	RefundTime    time.Time `json:"refund_time"`    // 退款时间
	RefundReason  string    `json:"refund_reason"`  // 退款原因
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// 创建支付请求
type CreatePaymentRequest struct {
	OrderID    string  `json:"order_id" validate:"required"`
	UserID     string  `json:"user_id" validate:"required"`
	MerchantID string  `json:"merchant_id" validate:"required"`
	Amount     float64 `json:"amount" validate:"required,gt=0"`
	Method     string  `json:"method" validate:"required,oneof=alipay wechat card balance"`
}

// 处理支付结果请求
type ProcessPaymentResultRequest struct {
	OrderID       string `json:"order_id" validate:"required"`
	TransactionID string `json:"transaction_id" validate:"required"`
	Status        string `json:"status" validate:"required,oneof=success failed"`
	PaymentTime   string `json:"payment_time,omitempty"` // 格式: "2025-05-14T12:00:00Z"
}

// 退款请求
type RefundRequest struct {
	OrderID string  `json:"order_id" validate:"required"`
	Reason  string  `json:"reason" validate:"required,max=200"`
	Amount  float64 `json:"amount" validate:"required,gt=0"`
}
