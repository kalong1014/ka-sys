package repository

import (
	"analytics-service/internal/domain"
	"context"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AnalyticsMySQLRepo struct {
	db *gorm.DB
}

func NewAnalyticsMySQLRepo(dsn string) (*AnalyticsMySQLRepo, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&domain.TrafficRecord{})
	if err != nil {
		return nil, err
	}

	return &AnalyticsMySQLRepo{db: db}, nil
}

// 记录流量
func (r *AnalyticsMySQLRepo) RecordTraffic(ctx context.Context, record *domain.TrafficRecord) error {
	return r.db.Create(record).Error
}

// 获取域名流量统计
func (r *AnalyticsMySQLRepo) GetDomainTraffic(ctx context.Context, domainID string, startTime, endTime time.Time) ([]*domain.TrafficRecord, error) {
	var records []*domain.TrafficRecord

	err := r.db.Where("domain_id = ? AND created_at BETWEEN ? AND ?",
		domainID, startTime, endTime).
		Order("created_at ASC").
		Find(&records).Error

	return records, err
}
