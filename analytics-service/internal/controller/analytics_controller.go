package controller

import (
	"analytics-service/internal/domain"
	"analytics-service/internal/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AnalyticsController struct {
	analyticsService *service.AnalyticsService
}

func NewAnalyticsController(as *service.AnalyticsService) *AnalyticsController {
	return &AnalyticsController{
		analyticsService: as,
	}
}

// 创建访问记录
func (c *AnalyticsController) CreateVisit(ctx *gin.Context) {
	var req domain.CreateVisitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	visit, err := c.analyticsService.CreateVisit(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "创建访问记录失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, visit)
}

// 获取订单统计
func (c *AnalyticsController) GetOrderStats(ctx *gin.Context) {
	var req domain.GetStatsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	stats, err := c.analyticsService.GetOrderStats(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取订单统计失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stats)
}

// 获取流量统计
func (c *AnalyticsController) GetTrafficStats(ctx *gin.Context) {
	var req domain.GetStatsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	stats, err := c.analyticsService.GetTrafficStats(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取流量统计失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stats)
}
