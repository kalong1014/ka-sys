package domain

import (
	"time"
)

// 卡密状态
const (
	CardKeyStatusAvailable = "available" // 可用
	CardKeyStatusUsed      = "used"      // 已使用
	CardKeyStatusExpired   = "expired"   // 已过期
	CardKeyStatusFrozen    = "frozen"    // 已冻结
)

// 商品模型
type Product struct {
	ID          string    `json:"id"`
	MerchantID  string    `json:"merchant_id"` // 所属商户ID
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`  // 库存数量
	Status      string    `json:"status"` // 状态: "active", "inactive"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// 卡密模型
type CardKey struct {
	ID         string    `json:"id"`
	ProductID  string    `json:"product_id"`  // 所属商品ID
	MerchantID string    `json:"merchant_id"` // 所属商户ID
	Content    string    `json:"content"`     // 卡密内容
	Status     string    `json:"status"`      // 状态: available, used, expired, frozen
	UsedBy     string    `json:"used_by"`     // 使用用户ID
	UsedAt     time.Time `json:"used_at"`     // 使用时间
	ExpireAt   time.Time `json:"expire_at"`   // 过期时间
	CreatedAt  time.Time `json:"created_at"`
}

// 创建商品请求
type CreateProductRequest struct {
	MerchantID  string  `json:"merchant_id" validate:"required"`
	Name        string  `json:"name" validate:"required,min=2,max=50"`
	Description string  `json:"description" validate:"max=500"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"required,gte=0"`
}

// 创建卡密请求
type GenerateCardKeysRequest struct {
	ProductID  string `json:"product_id" validate:"required"`
	MerchantID string `json:"merchant_id" validate:"required"`
	Count      int    `json:"count" validate:"required,gt=0,lt=1000"` // 生成数量
	ExpireDays int    `json:"expire_days" validate:"required,gt=0"`   // 有效期天数
}
