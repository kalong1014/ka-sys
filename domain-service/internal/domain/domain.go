package domain

import (
	"time"
)

// 域名状态
const (
	DomainStatusActive   = "active"   // 已激活
	DomainStatusPending  = "pending"  // 待审核
	DomainStatusRejected = "rejected" // 已拒绝
	DomainStatusExpired  = "expired"  // 已过期
)

// 域名记录
type Domain struct {
	ID         string    `json:"id"`
	MerchantID string    `json:"merchant_id"` // 所属商户ID
	DomainName string    `json:"domain_name"` // 域名
	PageID     string    `json:"page_id"`     // 绑定的页面ID
	Status     string    `json:"status"`      // 状态
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ExpireAt   time.Time `json:"expire_at"` // 过期时间
}

// 创建域名请求
type CreateDomainRequest struct {
	MerchantID string `json:"merchant_id" validate:"required"`
	DomainName string `json:"domain_name" validate:"required"`
	PageID     string `json:"page_id" validate:"required"`
}

// 更新域名请求
type UpdateDomainRequest struct {
	PageID   string `json:"page_id"`
	Status   string `json:"status"`
	ExpireAt string `json:"expire_at"` // 格式: "2025-05-14T12:00:00Z"
}
