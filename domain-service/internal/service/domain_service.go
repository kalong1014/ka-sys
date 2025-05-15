package service

import (
	"context"
	"domain-service/internal/domain"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

// 本地状态常量映射
const (
	statusActive  = "active"
	statusExpired = "expired"
)

type DomainService struct {
	domains map[string]domain.Domain // 域名存储
}

func NewDomainService() *DomainService {
	return &DomainService{
		domains: make(map[string]domain.Domain),
	}
}

// 创建域名绑定
func (s *DomainService) CreateDomain(ctx context.Context, req *domain.CreateDomainRequest) (*domain.Domain, error) {
	log.Printf("开始创建域名绑定: 商户ID=%s, 域名=%s", req.MerchantID, req.DomainName)

	// 校验域名格式
	if !s.isValidDomain(req.DomainName) {
		log.Println("域名格式不正确")
		return nil, errors.New("域名格式不正确")
	}

	// 检查域名是否已存在
	if s.isDomainExists(req.DomainName) {
		log.Println("域名已存在")
		return nil, errors.New("域名已存在")
	}

	// 生成域名记录ID
	domainID := uuid.New().String()

	// 创建域名对象
	domain := domain.Domain{
		ID:         domainID,
		MerchantID: req.MerchantID,
		DomainName: req.DomainName,
		PageID:     req.PageID,
		Status:     domain.DomainStatusPending,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		ExpireAt:   time.Now().AddDate(0, 1, 0), // 默认1个月后过期
	}

	// 保存域名记录
	s.domains[domainID] = domain
	log.Printf("域名绑定创建成功: ID=%s, 域名=%s", domainID, req.DomainName)

	return &domain, nil
}

// 更新域名信息
func (s *DomainService) UpdateDomain(ctx context.Context, domainID string, req *domain.UpdateDomainRequest) (*domain.Domain, error) {
	log.Printf("开始更新域名信息: ID=%s", domainID)

	// 检查域名记录是否存在
	domain, exists := s.domains[domainID]
	if !exists {
		log.Println("域名记录不存在")
		return nil, errors.New("域名记录不存在")
	}

	// 更新页面ID
	if req.PageID != "" {
		domain.PageID = req.PageID
	}

	// 更新状态
	if req.Status != "" {
		domain.Status = req.Status
	}

	// 更新过期时间
	if req.ExpireAt != "" {
		expireAt, err := time.Parse(time.RFC3339, req.ExpireAt)
		if err != nil {
			log.Printf("解析过期时间失败: %v", err)
			return nil, errors.New("过期时间格式不正确")
		}
		domain.ExpireAt = expireAt
	}

	domain.UpdatedAt = time.Now()
	s.domains[domainID] = domain
	log.Printf("域名信息更新成功: ID=%s", domainID)

	return &domain, nil
}

// 获取域名详情
func (s *DomainService) GetDomain(ctx context.Context, domainID string) (*domain.Domain, error) {
	log.Printf("获取域名详情: ID=%s", domainID)

	domain, exists := s.domains[domainID]
	if !exists {
		log.Println("域名记录不存在")
		return nil, errors.New("域名记录不存在")
	}

	// 检查是否过期
	if domain.Status == statusActive && time.Now().After(domain.ExpireAt) {
		domain.Status = statusExpired // 使用本地常量
		s.domains[domainID] = domain
		log.Printf("域名已过期: ID=%s", domainID)
	}

	return &domain, nil
}

// 删除域名
func (s *DomainService) DeleteDomain(ctx context.Context, domainID string) error {
	log.Printf("开始删除域名: ID=%s", domainID)

	if _, exists := s.domains[domainID]; !exists {
		log.Println("域名记录不存在")
		return errors.New("域名记录不存在")
	}

	delete(s.domains, domainID)
	log.Printf("域名删除成功: ID=%s", domainID)

	return nil
}

// 校验域名格式
func (s *DomainService) isValidDomain(domain string) bool {
	// 简化的域名格式校验
	if len(domain) < 3 || len(domain) > 253 {
		return false
	}

	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		return false
	}

	for _, part := range parts {
		if len(part) == 0 {
			return false
		}
		// 检查是否包含非法字符
		if !isAlphaNumeric(part) {
			return false
		}
	}

	return true
}

// 检查域名是否已存在
func (s *DomainService) isDomainExists(domainName string) bool {
	for _, domain := range s.domains {
		if domain.DomainName == domainName {
			return true
		}
	}
	return false
}

// 检查字符串是否只包含字母、数字和连字符
func isAlphaNumeric(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') ||
			r == '-') {
			return false
		}
	}
	return true
}
