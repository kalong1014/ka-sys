package repository

import (
	"context"
	"domain-service/internal/domain"
	"errors"

	"gorm.io/gorm"
)

type DomainRepository struct {
	db *gorm.DB
}

func NewDomainRepository(db *gorm.DB) *DomainRepository {
	return &DomainRepository{db: db}
}

// 创建域名记录
func (r *DomainRepository) Create(ctx context.Context, domain *domain.Domain) error {
	dbDomain := FromDomain(domain)
	result := r.db.WithContext(ctx).Create(dbDomain)
	return result.Error
}

// 查询域名记录（通过ID）
func (r *DomainRepository) GetByID(ctx context.Context, id string) (*domain.Domain, error) {
	var dbDomain DbDomain
	result := r.db.WithContext(ctx).First(&dbDomain, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("记录不存在")
		}
		return nil, result.Error
	}
	return dbDomain.ToDomain(), nil
}

// 查询域名记录（通过域名）
func (r *DomainRepository) GetByDomainName(ctx context.Context, name string) (*domain.Domain, error) {
	var dbDomain DbDomain
	result := r.db.WithContext(ctx).First(&dbDomain, "domain_name = ?", name)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("域名未注册")
		}
		return nil, result.Error
	}
	return dbDomain.ToDomain(), nil
}

// 更新域名记录
func (r *DomainRepository) Update(ctx context.Context, domain *domain.Domain) error {
	dbDomain := FromDomain(domain)
	result := r.db.WithContext(ctx).Save(dbDomain)
	return result.Error
}

// 删除域名记录
func (r *DomainRepository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&DbDomain{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return errors.New("记录不存在")
	}
	return result.Error
}
