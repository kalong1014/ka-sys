package controller

import (
	"context"
	"net/http"
	"user-service/internal/domain"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(us *service.UserService) *UserController {
	return &UserController{
		userService: us,
	}
}

// 注册接口
func (c *UserController) Register(ctx *gin.Context) {
	var req domain.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层注册逻辑
	resp, err := c.userService.Register(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// 登录接口
func (c *UserController) Login(ctx *gin.Context) {
	var req domain.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层登录逻辑
	resp, err := c.userService.Login(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "认证失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
