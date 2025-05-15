package domain

import (
	"time"
)

// 商户状态
const (
	MerchantActive    = "active"    // 已激活
	MerchantPending   = "pending"   // 待审核
	MerchantRejected  = "rejected"  // 已拒绝
	MerchantSuspended = "suspended" // 已暂停
)

// 商户信息
type Merchant struct {
	ID           string    `json:"id"`
	Name         string    `json:"name" validate:"required"`
	ContactName  string    `json:"contact_name" validate:"required"`
	ContactEmail string    `json:"contact_email" validate:"required,email"`
	ContactPhone string    `json:"contact_phone" validate:"required"`
	Status       string    `json:"status" validate:"required,oneof=active pending rejected suspended"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// 创建商户请求
type CreateMerchantRequest struct {
	Name         string `json:"name" validate:"required"`
	ContactName  string `json:"contact_name" validate:"required"`
	ContactEmail string `json:"contact_email" validate:"required,email"`
	ContactPhone string `json:"contact_phone" validate:"required"`
}

// 更新商户请求
type UpdateMerchantRequest struct {
	Name         string `json:"name"`
	ContactName  string `json:"contact_name"`
	ContactEmail string `json:"contact_email"`
	ContactPhone string `json:"contact_phone"`
	Status       string `json:"status" validate:"omitempty,oneof=active pending rejected suspended"`
}
