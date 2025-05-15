package controller

import (
	"context"
	"net/http"
	"order-service/internal/domain"
	"order-service/internal/service"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(os *service.OrderService) *OrderController {
	return &OrderController{
		orderService: os,
	}
}

// 创建订单
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var req domain.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	order, err := c.orderService.CreateOrder(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "创建订单失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// 支付订单
func (c *OrderController) PayOrder(ctx *gin.Context) {
	var req domain.PayOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	order, err := c.orderService.PayOrder(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "支付订单失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// 获取订单信息
func (c *OrderController) GetOrder(ctx *gin.Context) {
	orderID := ctx.Param("id")
	if orderID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "订单ID不能为空"})
		return
	}

	// 调用服务层逻辑
	order, err := c.orderService.GetOrder(context.Background(), orderID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "获取订单信息失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// 获取用户订单列表
func (c *OrderController) GetUserOrders(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户ID不能为空"})
		return
	}

	// 调用服务层逻辑
	orders, err := c.orderService.GetUserOrders(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户订单列表失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
