package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"product-service/internal/domain"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ProductService struct {
	products map[string]domain.Product // 商品存储
	cardKeys map[string]domain.CardKey // 卡密存储
}

func NewProductService() *ProductService {
	return &ProductService{
		products: make(map[string]domain.Product),
		cardKeys: make(map[string]domain.CardKey),
	}
}

// 创建商品
func (s *ProductService) CreateProduct(ctx context.Context, req *domain.CreateProductRequest) (*domain.Product, error) {
	log.Println("开始创建商品")

	// 检查商户是否存在（实际需调用商户服务验证）
	// 此处简化处理，假设商户ID有效

	// 生成商品ID
	productID := uuid.New().String()

	// 创建商品对象
	product := domain.Product{
		ID:          productID,
		MerchantID:  req.MerchantID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// 保存商品
	s.products[productID] = product
	log.Printf("商品创建成功: ID=%s, 名称=%s", productID, req.Name)

	return &product, nil
}

// 生成卡密
func (s *ProductService) GenerateCardKeys(ctx context.Context, req *domain.GenerateCardKeysRequest) (*[]domain.CardKey, error) {
	log.Printf("开始为商品 %s 生成 %d 个卡密", req.ProductID, req.Count)

	// 检查商品是否存在
	product, exists := s.products[req.ProductID]
	if !exists {
		log.Println("商品不存在")
		return nil, errors.New("商品不存在")
	}

	// 检查商户ID是否匹配
	if product.MerchantID != req.MerchantID {
		log.Println("商户ID不匹配")
		return nil, errors.New("商户ID不匹配")
	}

	// 计算过期时间
	expireAt := time.Now().AddDate(0, 0, req.ExpireDays)

	// 生成卡密
	var cardKeys []domain.CardKey
	for i := 0; i < req.Count; i++ {
		cardKeyID := uuid.New().String()
		content := s.generateCardKeyContent()

		cardKey := domain.CardKey{
			ID:         cardKeyID,
			ProductID:  req.ProductID,
			MerchantID: req.MerchantID,
			Content:    content,
			Status:     domain.CardKeyStatusAvailable,
			ExpireAt:   expireAt,
			CreatedAt:  time.Now(),
		}

		s.cardKeys[cardKeyID] = cardKey
		cardKeys = append(cardKeys, cardKey)
	}

	// 更新商品库存
	product.Stock += req.Count
	s.products[req.ProductID] = product

	log.Printf("成功生成 %d 个卡密", req.Count)
	return &cardKeys, nil
}

// 获取卡密信息
func (s *ProductService) GetCardKey(ctx context.Context, cardKeyID string) (*domain.CardKey, error) {
	log.Printf("获取卡密信息: ID=%s", cardKeyID)

	cardKey, exists := s.cardKeys[cardKeyID]
	if !exists {
		log.Println("卡密不存在")
		return nil, errors.New("卡密不存在")
	}

	// 检查是否过期
	if cardKey.Status == domain.CardKeyStatusAvailable && time.Now().After(cardKey.ExpireAt) {
		cardKey.Status = domain.CardKeyStatusExpired
		s.cardKeys[cardKeyID] = cardKey
	}

	return &cardKey, nil
}

// 更新卡密状态
func (s *ProductService) UpdateCardKeyStatus(ctx context.Context, cardKeyID, status, usedBy string) error {
	log.Printf("更新卡密状态: ID=%s, 新状态=%s", cardKeyID, status)

	cardKey, exists := s.cardKeys[cardKeyID]
	if !exists {
		log.Println("卡密不存在")
		return errors.New("卡密不存在")
	}

	// 检查状态转换是否合法
	if !s.isValidStatusTransition(cardKey.Status, status) {
		log.Printf("非法状态转换: 从 %s 到 %s", cardKey.Status, status)
		return errors.New("非法状态转换")
	}

	// 更新状态
	cardKey.Status = status
	if status == domain.CardKeyStatusUsed {
		cardKey.UsedBy = usedBy
		cardKey.UsedAt = time.Now()
	}
	s.cardKeys[cardKeyID] = cardKey

	log.Println("卡密状态更新成功")
	return nil
}

// 生成卡密内容
func (s *ProductService) generateCardKeyContent() string {
	// 生成随机字节
	bytes := make([]byte, 16)
	rand.Read(bytes)

	// 转换为Base64编码并移除特殊字符
	encoded := base64.RawURLEncoding.EncodeToString(bytes)

	// 确保长度为24个字符
	if len(encoded) > 24 {
		return encoded[:24]
	}

	// 补全到24个字符
	return encoded + strings.Repeat("A", 24-len(encoded))
}

// 检查状态转换是否合法
func (s *ProductService) isValidStatusTransition(oldStatus, newStatus string) bool {
	validTransitions := map[string][]string{
		domain.CardKeyStatusAvailable: {domain.CardKeyStatusUsed, domain.CardKeyStatusFrozen, domain.CardKeyStatusExpired},
		domain.CardKeyStatusFrozen:    {domain.CardKeyStatusAvailable},
		domain.CardKeyStatusExpired:   {},
		domain.CardKeyStatusUsed:      {},
	}

	for _, status := range validTransitions[oldStatus] {
		if status == newStatus {
			return true
		}
	}

	return false
}
