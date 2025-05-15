package domain

import (
	"time"
)

// 统计周期
const (
	PeriodDaily   = "daily"   // 日统计
	PeriodWeekly  = "weekly"  // 周统计
	PeriodMonthly = "monthly" // 月统计
	PeriodYearly  = "yearly"  // 年统计
)

// 访问记录
type VisitRecord struct {
	ID        string    `json:"id"`
	DomainID  string    `json:"domain_id"`  // 域名ID
	PageID    string    `json:"page_id"`    // 页面ID
	UserID    string    `json:"user_id"`    // 用户ID（匿名用户为""）
	IP        string    `json:"ip"`         // IP地址
	UserAgent string    `json:"user_agent"` // 用户代理
	Referrer  string    `json:"referrer"`   // 来源URL
	VisitTime time.Time `json:"visit_time"` // 访问时间
	Duration  int       `json:"duration"`   // 停留时长（秒）
}

// 订单统计
type OrderStats struct {
	Period       string    `json:"period"`        // 统计周期
	StartDate    time.Time `json:"start_date"`    // 开始日期
	EndDate      time.Time `json:"end_date"`      // 结束日期
	OrderCount   int       `json:"order_count"`   // 订单数量
	TotalAmount  float64   `json:"total_amount"`  // 总金额
	SuccessCount int       `json:"success_count"` // 成功订单数
	FailCount    int       `json:"fail_count"`    // 失败订单数
}

// 流量统计
type TrafficStats struct {
	Period       string    `json:"period"`       // 统计周期
	StartDate    time.Time `json:"start_date"`   // 开始日期
	EndDate      time.Time `json:"end_date"`     // 结束日期
	VisitCount   int       `json:"visit_count"`  // 访问次数
	UniqueUsers  int       `json:"unique_users"` // 独立用户数
	AvgDuration  int       `json:"avg_duration"` // 平均停留时长（秒）
	TopReferrers []struct {
		Referrer string `json:"referrer"`
		Count    int    `json:"count"`
	} `json:"top_referrers"` // 主要来源
}

// 创建访问记录请求
type CreateVisitRequest struct {
	DomainID  string `json:"domain_id" validate:"required"`
	PageID    string `json:"page_id" validate:"required"`
	UserID    string `json:"user_id"`
	IP        string `json:"ip" validate:"required"`
	UserAgent string `json:"user_agent" validate:"required"`
	Referrer  string `json:"referrer"`
	Duration  int    `json:"duration"`
}

// 获取统计数据请求
type GetStatsRequest struct {
	MerchantID string `json:"merchant_id" validate:"required"`
	DomainID   string `json:"domain_id"`
	PageID     string `json:"page_id"`
	Period     string `json:"period" validate:"required,oneof=daily weekly monthly yearly"`
	StartDate  string `json:"start_date"` // 格式: "2025-05-14"
	EndDate    string `json:"end_date"`   // 格式: "2025-05-15"
}
