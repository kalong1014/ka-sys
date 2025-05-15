package service

import (
	"context"
	"errors"
	"log"
	"merchant-service/internal/domain"
	"time"
	"merchant-service/internal/repository"
	"github.com/google/uuid"
)

type MerchantService struct {
	applications map[string]domain.MerchantApplication
	merchants    map[string]domain.Merchant
	repo repository.MerchantRepository // 修改接口类型
}

func NewMerchantService(repo repository.MerchantRepository) *MerchantService {
	return &MerchantService{
		repo: repo
		applications: make(map[string]domain.MerchantApplication),
		merchants:    make(map[string]domain.Merchant),
	}
}

// 提交入驻申请
func (s *MerchantService) ApplyMerchant(ctx context.Context, req *domain.MerchantApplyRequest) (*domain.MerchantApplyResponse, error) {
	log.Println("开始处理商户入驻申请请求")
	// 检查用户是否已提交申请
	for _, app := range s.applications {
		if app.UserID == req.UserID && app.Status == domain.MerchantStatusPending {
			log.Println("已有待审核的申请")
			return nil, errors.New("已有待审核的申请")
		}
	}

	// 检查用户是否已是商户
	for _, merchant := range s.merchants {
		if merchant.UserID == req.UserID {
			log.Println("用户已成为商户")
			return nil, errors.New("用户已成为商户")
		}
	}

	// 创建申请ID
	appID := uuid.New().String()

	// 创建入驻申请对象
	application := domain.MerchantApplication{
		ID:              appID,
		UserID:          req.UserID,
		Name:            req.Name,
		Description:     req.Description,
		BusinessLicense: req.BusinessLicense,
		ContactName:     req.ContactName,
		ContactPhone:    req.ContactPhone,
		Status:          domain.MerchantStatusPending,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// 保存申请
	s.applications[appID] = application
	log.Println("商户入驻申请提交成功")

	return &domain.MerchantApplyResponse{
		ID:        appID,
		UserID:    req.UserID,
		Name:      req.Name,
		Status:    domain.MerchantStatusPending,
		CreatedAt: time.Now(),
	}, nil
}

// 审核商户申请
func (s *MerchantService) ReviewMerchant(ctx context.Context, req *domain.MerchantReviewRequest) error {
	log.Println("开始审核商户申请")
	// 检查申请是否存在
	app, exists := s.applications[req.ID]
	if !exists {
		log.Println("申请不存在")
		return errors.New("申请不存在")
	}

	// 检查状态是否已处理
	if app.Status != domain.MerchantStatusPending {
		log.Println("申请已处理")
		return errors.New("申请已处理")
	}

	// 更新申请状态
	app.Status = req.Status
	app.Reason = req.Reason
	app.UpdatedAt = time.Now()
	s.applications[req.ID] = app
	log.Println("申请状态更新成功")

	// 如果审核通过，创建商户记录
	if req.Status == domain.MerchantStatusApproved {
		merchant := domain.Merchant{
			ID:              app.ID,
			UserID:          app.UserID,
			Name:            app.Name,
			Description:     app.Description,
			BusinessLicense: app.BusinessLicense,
			ContactName:     app.ContactName,
			ContactPhone:    app.ContactPhone,
			Level:           1, // 默认基础等级
			Status:          domain.MerchantStatusApproved,
			CreatedAt:       app.CreatedAt,
			UpdatedAt:       time.Now(),
		}
		s.merchants[app.ID] = merchant
		log.Println("商户创建成功")
	}

	return nil
}

// 获取商户信息
func (s *MerchantService) GetMerchantInfo(ctx context.Context, merchantID string) (*domain.Merchant, error) {
	log.Println("开始获取商户信息")
	merchant, exists := s.merchants[merchantID]
	if !exists {
		log.Println("商户不存在")
		return nil, errors.New("商户不存在")
	}

	return &merchant, nil
}
