package service

import (
	"analytics-service/internal/domain"
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

type AnalyticsService struct {
	visits  map[string]domain.VisitRecord  // 访问记录存储
	orders  map[string]domain.OrderStats   // 订单统计存储
	traffic map[string]domain.TrafficStats // 流量统计存储
}

func NewAnalyticsService() *AnalyticsService {
	return &AnalyticsService{
		visits:  make(map[string]domain.VisitRecord),
		orders:  make(map[string]domain.OrderStats),
		traffic: make(map[string]domain.TrafficStats),
	}
}

// 创建访问记录
func (s *AnalyticsService) CreateVisit(ctx context.Context, req *domain.CreateVisitRequest) (*domain.VisitRecord, error) {
	log.Printf("开始记录访问: 域名ID=%s, 页面ID=%s, IP=%s",
		req.DomainID, req.PageID, req.IP)

	// 生成访问记录ID
	visitID := uuid.New().String()

	// 创建访问记录
	visit := domain.VisitRecord{
		ID:        visitID,
		DomainID:  req.DomainID,
		PageID:    req.PageID,
		UserID:    req.UserID,
		IP:        req.IP,
		UserAgent: req.UserAgent,
		Referrer:  req.Referrer,
		VisitTime: time.Now(),
		Duration:  req.Duration,
	}

	// 保存访问记录
	s.visits[visitID] = visit
	log.Printf("访问记录成功: ID=%s", visitID)

	// 更新流量统计（简化处理，实际应异步处理）
	s.updateTrafficStats(ctx, &visit)

	return &visit, nil
}

// 获取订单统计
func (s *AnalyticsService) GetOrderStats(ctx context.Context, req *domain.GetStatsRequest) (*domain.OrderStats, error) {
	log.Printf("获取订单统计: 商户ID=%s, 周期=%s", req.MerchantID, req.Period)

	// 解析日期范围
	startDate, endDate, err := s.parseDateRange(req.Period, req.StartDate, req.EndDate)
	if err != nil {
		log.Printf("解析日期范围失败: %v", err)
		return nil, err
	}

	// 生成统计ID
	statsID := s.generateStatsID(req.MerchantID, req.DomainID, req.PageID, req.Period, startDate, endDate)

	// 检查缓存（实际应从数据库或缓存中获取）
	if stats, exists := s.orders[statsID]; exists {
		log.Printf("从缓存获取订单统计: ID=%s", statsID)
		return &stats, nil
	}

	// 模拟统计计算（实际应从订单服务获取数据）
	stats := domain.OrderStats{
		Period:       req.Period,
		StartDate:    startDate,
		EndDate:      endDate,
		OrderCount:   100,
		TotalAmount:  9999.99,
		SuccessCount: 85,
		FailCount:    15,
	}

	// 保存统计结果
	s.orders[statsID] = stats
	log.Printf("计算并保存订单统计: ID=%s", statsID)

	return &stats, nil
}

// 获取流量统计
func (s *AnalyticsService) GetTrafficStats(ctx context.Context, req *domain.GetStatsRequest) (*domain.TrafficStats, error) {
	log.Printf("获取流量统计: 商户ID=%s, 周期=%s", req.MerchantID, req.Period)

	// 解析日期范围
	startDate, endDate, err := s.parseDateRange(req.Period, req.StartDate, req.EndDate)
	if err != nil {
		log.Printf("解析日期范围失败: %v", err)
		return nil, err
	}

	// 生成统计ID
	statsID := s.generateStatsID(req.MerchantID, req.DomainID, req.PageID, req.Period, startDate, endDate)

	// 检查缓存
	if stats, exists := s.traffic[statsID]; exists {
		log.Printf("从缓存获取流量统计: ID=%s", statsID)
		return &stats, nil
	}

	// 模拟统计计算
	stats := domain.TrafficStats{
		Period:      req.Period,
		StartDate:   startDate,
		EndDate:     endDate,
		VisitCount:  1000,
		UniqueUsers: 500,
		AvgDuration: 180,
		TopReferrers: []struct {
			Referrer string `json:"referrer"`
			Count    int    `json:"count"`
		}{
			{"https://google.com", 400},
			{"https://facebook.com", 200},
			{"https://twitter.com", 150},
		},
	}

	// 保存统计结果
	s.traffic[statsID] = stats
	log.Printf("计算并保存流量统计: ID=%s", statsID)

	return &stats, nil
}

// 更新流量统计
func (s *AnalyticsService) updateTrafficStats(ctx context.Context, visit *domain.VisitRecord) {
	log.Printf("更新流量统计: 域名ID=%s, 页面ID=%s", visit.DomainID, visit.PageID)

	// 简化处理，实际应按天/周/月聚合
	// 此处仅记录访问，统计计算应由定时任务完成
}

// 解析日期范围
func (s *AnalyticsService) parseDateRange(period, startDateStr, endDateStr string) (time.Time, time.Time, error) {
	now := time.Now()
	var startDate, endDate time.Time

	switch period {
	case domain.PeriodDaily:
		startDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		endDate = startDate.AddDate(0, 0, 1)
	case domain.PeriodWeekly:
		weekday := int(now.Weekday())
		if weekday == 0 { // Sunday
			weekday = 7
		}
		startDate = time.Date(now.Year(), now.Month(), now.Day()-weekday+1, 0, 0, 0, 0, time.UTC)
		endDate = startDate.AddDate(0, 0, 7)
	case domain.PeriodMonthly:
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
		endDate = startDate.AddDate(0, 1, 0)
	case domain.PeriodYearly:
		startDate = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
		endDate = startDate.AddDate(1, 0, 0)
	default:
		return time.Time{}, time.Time{}, errors.New("不支持的统计周期")
	}

	// 如果提供了自定义日期范围，则使用自定义范围
	if startDateStr != "" {
		var err error
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return time.Time{}, time.Time{}, errors.New("开始日期格式不正确")
		}
	}

	if endDateStr != "" {
		var err error
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return time.Time{}, time.Time{}, errors.New("结束日期格式不正确")
		}
		endDate = endDate.AddDate(0, 0, 1) // 包含结束日期当天
	}

	return startDate, endDate, nil
}

// 生成统计ID
func (s *AnalyticsService) generateStatsID(merchantID, domainID, pageID, period string, startDate, endDate time.Time) string {
	parts := []string{
		"stats",
		merchantID,
		domainID,
		pageID,
		period,
		startDate.Format("20060102"),
		endDate.Format("20060102"),
	}

	return strings.Join(parts, "_")
}
