package controller

import (
	"context"
	"merchant-service/internal/domain"
	"merchant-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	merchantService *service.MerchantService
}

func NewMerchantController(ms *service.MerchantService) *MerchantController {
	return &MerchantController{
		merchantService: ms,
	}
}

// 提交入驻申请
func (c *MerchantController) ApplyMerchant(ctx *gin.Context) {
	var req domain.MerchantApplyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	resp, err := c.merchantService.ApplyMerchant(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "入驻申请失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// 审核商户申请（管理员接口）
func (c *MerchantController) ReviewMerchant(ctx *gin.Context) {
	var req domain.MerchantReviewRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	err := c.merchantService.ReviewMerchant(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "审核失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "审核完成"})
}

// 获取商户信息
func (c *MerchantController) GetMerchantInfo(ctx *gin.Context) {
	merchantID := ctx.Param("id")
	if merchantID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "商户ID不能为空"})
		return
	}

	// 调用服务层逻辑
	merchant, err := c.merchantService.GetMerchantInfo(context.Background(), merchantID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "获取商户信息失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, merchant)
}
