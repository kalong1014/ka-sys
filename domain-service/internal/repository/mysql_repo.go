package repository

import (
	"context"
	"domain-service/internal/domain"
	"errors"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DomainMySQLRepo struct {
	db *gorm.DB
}

func NewDomainMySQLRepo(dsn string) (*DomainMySQLRepo, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("连接数据库失败: %v", err)
		return nil, err
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&domain.Domain{})
	if err != nil {
		log.Printf("迁移表结构失败: %v", err)
		return nil, err
	}

	return &DomainMySQLRepo{db: db}, nil
}

// 创建域名
func (r *DomainMySQLRepo) CreateDomain(ctx context.Context, domain *domain.Domain) error {
	return r.db.Create(domain).Error
}

// 获取域名
func (r *DomainMySQLRepo) GetDomain(ctx context.Context, id string) (*domain.Domain, error) {
	var domain domain.Domain
	err := r.db.First(&domain, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &domain, err
}

// 更新域名
func (r *DomainMySQLRepo) UpdateDomain(ctx context.Context, domain *domain.Domain) error {
	return r.db.Save(domain).Error
}

// 删除域名
func (r *DomainMySQLRepo) DeleteDomain(ctx context.Context, id string) error {
	return r.db.Delete(&domain.Domain{}, "id = ?", id).Error
}

// 列出域名
func (r *DomainMySQLRepo) ListDomains(ctx context.Context, merchantID string) ([]*domain.Domain, error) {
	var domains []*domain.Domain
	err := r.db.Where("merchant_id = ?", merchantID).Find(&domains).Error
	return domains, err
}
