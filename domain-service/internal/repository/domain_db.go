package repository

import (
	"domain-service/internal/domain"
	"time"
)

// 数据库表模型（与domain.Domain字段映射）
type DbDomain struct {
	ID         string    `gorm:"primaryKey;column:id;type:varchar(36)"`
	MerchantID string    `gorm:"column:merchant_id;type:varchar(36);not null"`
	DomainName string    `gorm:"column:domain_name;type:varchar(255);uniqueIndex;not null"`
	PageID     string    `gorm:"column:page_id;type:varchar(36)"`
	Status     string    `gorm:"column:status;type:varchar(20);not null"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;not null"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;not null"`
	ExpireAt   time.Time `gorm:"column:expire_at;type:datetime"`
}

// 模型转换：DbDomain -> domain.Domain
func (d *DbDomain) ToDomain() *domain.Domain {
	return &domain.Domain{
		ID:         d.ID,
		MerchantID: d.MerchantID,
		DomainName: d.DomainName,
		PageID:     d.PageID,
		Status:     d.Status,
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
		ExpireAt:   d.ExpireAt,
	}
}

// 模型转换：domain.Domain -> DbDomain
func FromDomain(d *domain.Domain) *DbDomain {
	return &DbDomain{
		ID:         d.ID,
		MerchantID: d.MerchantID,
		DomainName: d.DomainName,
		PageID:     d.PageID,
		Status:     d.Status,
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
		ExpireAt:   d.ExpireAt,
	}
}
