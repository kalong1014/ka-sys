package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"page-builder-service/internal/domain"
	"sort"
	"time"

	"github.com/google/uuid"
)

type PageService struct {
	pages map[string]domain.Page // 页面存储
}

func NewPageService() *PageService {
	return &PageService{
		pages: make(map[string]domain.Page),
	}
}

// 创建页面
func (s *PageService) CreatePage(ctx context.Context, req *domain.CreatePageRequest) (*domain.Page, error) {
	log.Printf("开始创建页面: 商户ID=%s, 页面名称=%s", req.MerchantID, req.Name)

	// 校验元素合法性
	if err := s.validateElements(req.Elements); err != nil {
		log.Printf("页面元素校验失败: %v", err)
		return nil, err
	}

	// 生成页面ID
	pageID := uuid.New().String()

	// 创建页面对象
	page := domain.Page{
		ID:          pageID,
		MerchantID:  req.MerchantID,
		Name:        req.Name,
		Description: req.Description,
		Elements:    req.Elements,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// 保存页面
	s.pages[pageID] = page
	log.Printf("页面创建成功: ID=%s, 名称=%s", pageID, req.Name)

	return &page, nil
}

// 更新页面
func (s *PageService) UpdatePage(ctx context.Context, pageID string, req *domain.UpdatePageRequest) (*domain.Page, error) {
	log.Printf("开始更新页面: ID=%s", pageID)

	// 检查页面是否存在
	page, exists := s.pages[pageID]
	if !exists {
		log.Println("页面不存在")
		return nil, errors.New("页面不存在")
	}

	// 校验元素合法性
	if err := s.validateElements(req.Elements); err != nil {
		log.Printf("页面元素校验失败: %v", err)
		return nil, err
	}

	// 更新元素并排序
	page.Elements = req.Elements
	sort.Slice(page.Elements, func(i, j int) bool {
		// 按元素Position排序（假设所有元素都包含Position字段）
		elemI, _ := page.Elements[i].(map[string]interface{})
		elemJ, _ := page.Elements[j].(map[string]interface{})
		posI, _ := elemI["position"].(float64)
		posJ, _ := elemJ["position"].(float64)
		return int(posI) < int(posJ)
	})

	page.UpdatedAt = time.Now()
	s.pages[pageID] = page
	log.Printf("页面更新成功: ID=%s", pageID)

	return &page, nil
}

// 获取页面详情
func (s *PageService) GetPage(ctx context.Context, pageID string) (*domain.Page, error) {
	log.Printf("获取页面详情: ID=%s", pageID)

	page, exists := s.pages[pageID]
	if !exists {
		log.Println("页面不存在")
		return nil, errors.New("页面不存在")
	}

	return &page, nil
}

// 删除页面
func (s *PageService) DeletePage(ctx context.Context, pageID string) error {
	log.Printf("开始删除页面: ID=%s", pageID)

	if _, exists := s.pages[pageID]; !exists {
		log.Println("页面不存在")
		return errors.New("页面不存在")
	}

	delete(s.pages, pageID)
	log.Printf("页面删除成功: ID=%s", pageID)

	return nil
}

// 校验元素合法性
func (s *PageService) validateElements(elements []interface{}) error {
	allowedTypes := map[string]bool{
		domain.ElementTypeTitle:  true,
		domain.ElementTypeText:   true,
		domain.ElementTypeImage:  true,
		domain.ElementTypeButton: true,
		domain.ElementTypeList:   true,
	}

	for _, elem := range elements {
		elemMap, ok := elem.(map[string]interface{})
		if !ok {
			return errors.New("元素格式不正确，必须为对象")
		}

		elemType, exists := elemMap["type"].(string)
		if !exists {
			return errors.New("元素缺少type字段")
		}

		if !allowedTypes[elemType] {
			return fmt.Errorf("不支持的元素类型: %s", elemType)
		}

		// 校验必填字段
		switch elemType {
		case domain.ElementTypeTitle:
			if _, exists := elemMap["content"]; !exists {
				return errors.New("标题元素缺少content字段")
			}
			if _, exists := elemMap["level"]; !exists {
				return errors.New("标题元素缺少level字段")
			}
		case domain.ElementTypeText:
			if _, exists := elemMap["content"]; !exists {
				return errors.New("文本元素缺少content字段")
			}
		case domain.ElementTypeImage:
			if _, exists := elemMap["url"]; !exists {
				return errors.New("图片元素缺少url字段")
			}
		case domain.ElementTypeButton:
			if _, exists := elemMap["content"]; !exists {
				return errors.New("按钮元素缺少content字段")
			}
		case domain.ElementTypeList:
			if _, exists := elemMap["items"]; !exists {
				return errors.New("列表元素缺少items字段")
			}
		}

		// 校验Position字段（必须为整数且≥0）
		pos, exists := elemMap["position"]
		if !exists {
			return errors.New("元素缺少position字段")
		}
		posFloat, ok := pos.(float64)
		if !ok || int(posFloat) < 0 {
			return errors.New("position字段必须为非负整数")
		}
	}

	return nil
}
