package repository

import (
	"context"
	"errors"
	"log"
	"merchant-service/internal/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MerchantMySQLRepo struct {
	db *gorm.DB
}

func NewMerchantMySQLRepo(dsn string) (*MerchantMySQLRepo, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("连接数据库失败: %v", err)
		return nil, err
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&domain.Merchant{})
	if err != nil {
		log.Printf("迁移表结构失败: %v", err)
		return nil, err
	}

	return &MerchantMySQLRepo{db: db}, nil
}

// 创建商户
func (r *MerchantMySQLRepo) CreateMerchant(ctx context.Context, merchant *domain.Merchant) error {
	return r.db.Create(merchant).Error
}

// 获取商户
func (r *MerchantMySQLRepo) GetMerchant(ctx context.Context, id string) (*domain.Merchant, error) {
	var merchant domain.Merchant
	err := r.db.First(&merchant, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &merchant, err
}

// 更新商户
func (r *MerchantMySQLRepo) UpdateMerchant(ctx context.Context, merchant *domain.Merchant) error {
	return r.db.Save(merchant).Error
}

// 删除商户
func (r *MerchantMySQLRepo) DeleteMerchant(ctx context.Context, id string) error {
	return r.db.Delete(&domain.Merchant{}, "id = ?", id).Error
}

// 列出商户
func (r *MerchantMySQLRepo) ListMerchants(ctx context.Context, page, pageSize int) ([]*domain.Merchant, error) {
	var merchants []*domain.Merchant
	offset := (page - 1) * pageSize

	err := r.db.Offset(offset).Limit(pageSize).Find(&merchants).Error
	return merchants, err
}
