package repository

import (
	"context"
	"errors"
	"order-service/internal/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type OrderMySQLRepo struct {
	db *gorm.DB
}

func NewOrderMySQLRepo(dsn string) (*OrderMySQLRepo, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&domain.Order{})
	if err != nil {
		return nil, err
	}

	return &OrderMySQLRepo{db: db}, nil
}

// 创建订单
func (r *OrderMySQLRepo) CreateOrder(ctx context.Context, order *domain.Order) error {
	return r.db.Create(order).Error
}

// 获取订单
func (r *OrderMySQLRepo) GetOrder(ctx context.Context, id string) (*domain.Order, error) {
	var order domain.Order
	err := r.db.First(&order, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &order, err
}

// 更新订单
func (r *OrderMySQLRepo) UpdateOrder(ctx context.Context, order *domain.Order) error {
	return r.db.Save(order).Error
}

// 列出商户订单
func (r *OrderMySQLRepo) ListOrders(ctx context.Context, merchantID string, page, pageSize int) ([]*domain.Order, error) {
	var orders []*domain.Order
	offset := (page - 1) * pageSize

	err := r.db.Where("merchant_id = ?", merchantID).
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&orders).Error

	return orders, err
}
