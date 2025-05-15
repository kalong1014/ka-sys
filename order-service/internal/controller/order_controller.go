package controller

import (
	"context"
	"net/http"
	"order-service/internal/domain"
	"order-service/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(os *service.OrderService) *OrderController {
	return &OrderController{orderService: os}
}

// 创建订单
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var req domain.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数解析错误", "detail": err.Error()})
		return
	}

	order, err := c.orderService.CreateOrder(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "创建订单失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

// 获取订单
func (c *OrderController) GetOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	order, err := c.orderService.GetOrder(context.Background(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取订单失败", "detail": err.Error()})
		return
	}

	if order == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "订单不存在"})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// 列出商户订单
func (c *OrderController) ListOrders(ctx *gin.Context) {
	merchantID := ctx.Param("merchant_id")
	pageStr := ctx.DefaultQuery("page", "1")
	pageSizeStr := ctx.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	orders, err := c.orderService.ListOrders(context.Background(), merchantID, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取订单列表失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
