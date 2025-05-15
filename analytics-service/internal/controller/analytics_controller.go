package controller

import (
	"analytics-service/internal/service"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AnalyticsController struct {
	analyticsService *service.AnalyticsService
}

func NewAnalyticsController(as *service.AnalyticsService) *AnalyticsController {
	return &AnalyticsController{analyticsService: as}
}

// 获取商户概览数据
func (c *AnalyticsController) GetMerchantOverview(ctx *gin.Context) {
	merchantID := ctx.Param("merchant_id")
	today := time.Now()
	lastMonth := today.AddDate(0, -1, 0)

	overview, err := c.analyticsService.GetMerchantOverview(
		context.Background(),
		merchantID,
		lastMonth,
		today,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取概览数据失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, overview)
}

// 获取域名流量数据
func (c *AnalyticsController) GetDomainTraffic(ctx *gin.Context) {
	domainID := ctx.Param("domain_id")
	daysStr := ctx.DefaultQuery("days", "7")

	days, err := time.ParseDuration(daysStr + "d")
	if err != nil {
		days = 7 * 24 * time.Hour // 默认7天
	}

	endTime := time.Now()
	startTime := endTime.Add(-days)

	traffic, err := c.analyticsService.GetDomainTraffic(
		context.Background(),
		domainID,
		startTime,
		endTime,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取流量数据失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, traffic)
}
