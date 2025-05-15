package controller

import (
	"context"
	"net/http"
	"payment-service/internal/domain"
	"payment-service/internal/service"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	paymentService *service.PaymentService
}

func NewPaymentController(ps *service.PaymentService) *PaymentController {
	return &PaymentController{
		paymentService: ps,
	}
}

// 创建支付
func (c *PaymentController) CreatePayment(ctx *gin.Context) {
	var req domain.CreatePaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	payment, err := c.paymentService.CreatePayment(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "创建支付失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payment)
}

// 处理支付结果
func (c *PaymentController) ProcessPaymentResult(ctx *gin.Context) {
	var req domain.ProcessPaymentResultRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	payment, err := c.paymentService.ProcessPaymentResult(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "处理支付结果失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payment)
}

// 处理退款
func (c *PaymentController) ProcessRefund(ctx *gin.Context) {
	var req domain.RefundRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层逻辑
	payment, err := c.paymentService.ProcessRefund(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "处理退款失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payment)
}

// 获取支付信息
func (c *PaymentController) GetPayment(ctx *gin.Context) {
	paymentID := ctx.Param("id")
	if paymentID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "支付ID不能为空"})
		return
	}

	// 调用服务层逻辑
	payment, err := c.paymentService.GetPayment(context.Background(), paymentID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "获取支付信息失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payment)
}

// 获取订单的支付记录
func (c *PaymentController) GetOrderPayments(ctx *gin.Context) {
	orderID := ctx.Param("order_id")
	if orderID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "订单ID不能为空"})
		return
	}

	// 调用服务层逻辑
	payments, err := c.paymentService.GetOrderPayments(context.Background(), orderID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取订单支付记录失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payments)
}
