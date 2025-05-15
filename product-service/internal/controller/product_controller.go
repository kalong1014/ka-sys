package controller

import (
	"context"
	"net/http"
	"product-service/internal/domain"
	"product-service/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(ps *service.ProductService) *ProductController {
	return &ProductController{
		productService: ps,
	}
}

// 创建商品
func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var req domain.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	product, err := c.productService.CreateProduct(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "创建商品失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// 生成卡密
func (c *ProductController) GenerateCardKeys(ctx *gin.Context) {
	var req domain.GenerateCardKeysRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	cardKeys, err := c.productService.GenerateCardKeys(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "生成卡密失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cardKeys)
}

// 获取卡密信息
func (c *ProductController) GetCardKey(ctx *gin.Context) {
	cardKeyID := ctx.Param("id")
	if cardKeyID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "卡密ID不能为空"})
		return
	}

	// 调用服务层逻辑
	cardKey, err := c.productService.GetCardKey(context.Background(), cardKeyID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "获取卡密信息失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cardKey)
}

// 更新卡密状态
func (c *ProductController) UpdateCardKeyStatus(ctx *gin.Context) {
	cardKeyID := ctx.Param("id")
	if cardKeyID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "卡密ID不能为空"})
		return
	}

	status := ctx.Query("status")
	if status == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "状态不能为空"})
		return
	}

	usedBy := ctx.Query("used_by")

	// 调用服务层逻辑
	err := c.productService.UpdateCardKeyStatus(context.Background(), cardKeyID, status, usedBy)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新卡密状态失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "卡密状态更新成功"})
}
