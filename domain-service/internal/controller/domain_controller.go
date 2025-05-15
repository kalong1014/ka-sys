package controller

import (
	"context"
	"domain-service/internal/domain"
	"domain-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DomainController struct {
	domainService *service.DomainService
}

func NewDomainController(ds *service.DomainService) *DomainController {
	return &DomainController{
		domainService: ds,
	}
}

// 创建域名绑定
func (c *DomainController) CreateDomain(ctx *gin.Context) {
	var req domain.CreateDomainRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	domain, err := c.domainService.CreateDomain(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "创建域名绑定失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain)
}

// 更新域名信息
func (c *DomainController) UpdateDomain(ctx *gin.Context) {
	domainID := ctx.Param("id")
	if domainID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "域名ID不能为空"})
		return
	}

	var req domain.UpdateDomainRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	updatedDomain, err := c.domainService.UpdateDomain(context.Background(), domainID, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新域名信息失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedDomain)
}

// 获取域名详情
func (c *DomainController) GetDomain(ctx *gin.Context) {
	domainID := ctx.Param("id")
	if domainID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "域名ID不能为空"})
		return
	}

	// 调用服务层逻辑
	domain, err := c.domainService.GetDomain(context.Background(), domainID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "获取域名信息失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain)
}

// 删除域名
func (c *DomainController) DeleteDomain(ctx *gin.Context) {
	domainID := ctx.Param("id")
	if domainID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "域名ID不能为空"})
		return
	}

	// 调用服务层逻辑
	err := c.domainService.DeleteDomain(context.Background(), domainID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "删除域名失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "域名删除成功"})
}
