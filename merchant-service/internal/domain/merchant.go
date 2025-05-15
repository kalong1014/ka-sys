package domain

import (
	"time"
)

// 商户入驻申请状态
const (
	MerchantStatusPending  = "pending"  // 待审核
	MerchantStatusApproved = "approved" // 已通过
	MerchantStatusRejected = "rejected" // 已拒绝
)

// 商户入驻申请
type MerchantApplication struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	BusinessLicense string    `json:"business_license"`
	ContactName     string    `json:"contact_name"`
	ContactPhone    string    `json:"contact_phone"`
	Status          string    `json:"status"`
	Reason          string    `json:"reason,omitempty"` // 拒绝原因
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// 商户信息
type Merchant struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	BusinessLicense string    `json:"business_license"`
	ContactName     string    `json:"contact_name"`
	ContactPhone    string    `json:"contact_phone"`
	Level           int       `json:"level"` // 商户等级: 1=基础, 2=VIP, 3=旗舰
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// 入驻申请请求
type MerchantApplyRequest struct {
	UserID          string `json:"user_id" validate:"required"`
	Name            string `json:"name" validate:"required,min=2,max=50"`
	Description     string `json:"description" validate:"max=500"`
	BusinessLicense string `json:"business_license" validate:"required,min=10,max=50"`
	ContactName     string `json:"contact_name" validate:"required,min=2,max=20"`
	ContactPhone    string `json:"contact_phone" validate:"required,len=11"`
}

// 入驻申请响应
type MerchantApplyResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// 审核商户请求
type MerchantReviewRequest struct {
	ID     string `json:"id" validate:"required"`
	Status string `json:"status" validate:"required,oneof=approved rejected"`
	Reason string `json:"reason,omitempty"` // 拒绝时必填
}
