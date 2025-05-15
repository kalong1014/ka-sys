package events

// 订单创建事件
type OrderCreatedEvent struct {
	OrderID     string  `json:"order_id"`
	MerchantID  string  `json:"merchant_id"`
	TotalAmount float64 `json:"total_amount"`
}

// 支付成功事件
type PaymentSucceededEvent struct {
	OrderID   string `json:"order_id"`
	PaymentID string `json:"payment_id"`
}

// 支付失败事件
type PaymentFailedEvent struct {
	OrderID string `json:"order_id"`
	Reason  string `json:"reason"`
}
