package controller

import (
	"context"
	"net/http"
	"page-builder-service/internal/domain"
	"page-builder-service/internal/service"

	"github.com/gin-gonic/gin"
)

type PageController struct {
	pageService *service.PageService
}

func NewPageController(ps *service.PageService) *PageController {
	return &PageController{
		pageService: ps,
	}
}

// 创建页面
func (c *PageController) CreatePage(ctx *gin.Context) {
	var req domain.CreatePageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	page, err := c.pageService.CreatePage(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "创建页面失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, page)
}

// 更新页面
func (c *PageController) UpdatePage(ctx *gin.Context) {
	pageID := ctx.Param("id")
	if pageID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "页面ID不能为空"})
		return
	}

	var req domain.UpdatePageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	updatedPage, err := c.pageService.UpdatePage(context.Background(), pageID, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新页面失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedPage)
}

// 获取页面详情
func (c *PageController) GetPage(ctx *gin.Context) {
	pageID := ctx.Param("id")
	if pageID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "页面ID不能为空"})
		return
	}

	// 调用服务层逻辑
	page, err := c.pageService.GetPage(context.Background(), pageID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "获取页面失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, page)
}

// 删除页面
func (c *PageController) DeletePage(ctx *gin.Context) {
	pageID := ctx.Param("id")
	if pageID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "页面ID不能为空"})
		return
	}

	// 调用服务层逻辑
	err := c.pageService.DeletePage(context.Background(), pageID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "删除页面失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "页面删除成功"})
}
