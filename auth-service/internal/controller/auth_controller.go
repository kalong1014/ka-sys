package controller

import (
	"auth-service/internal/domain"
	"auth-service/internal/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	authService *service.AuthService
	validator   *validator.Validate
}

func NewAuthController(as *service.AuthService) *AuthController {
	return &AuthController{
		authService: as,
		validator:   validator.New(),
	}
}

// 用户注册
func (c *AuthController) Register(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数解析错误", "detail": err.Error()})
		return
	}

	// 校验参数
	if err := c.validator.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数校验失败", "detail": err.Error()})
		return
	}

	// 设置默认角色为普通用户
	if user.Role == "" {
		user.Role = domain.RoleUser
	}

	// 调用服务层
	err := c.authService.RegisterUser(context.Background(), &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
}

// 用户登录
func (c *AuthController) Login(ctx *gin.Context) {
	var req domain.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数解析错误", "detail": err.Error()})
		return
	}

	// 调用服务层
	tokenResp, err := c.authService.Login(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "登录失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tokenResp)
}

// 验证令牌（供网关调用）
func (c *AuthController) ValidateToken(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "缺少令牌"})
		return
	}

	claims, err := c.authService.ValidateToken(tokenStr)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "无效令牌", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, claims)
}

// 检查权限
func (c *AuthController) CheckPermission(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("Authorization")
	requiredRole := ctx.Query("role")

	if tokenStr == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "缺少令牌"})
		return
	}

	if requiredRole == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "缺少角色参数"})
		return
	}

	hasPermission, err := c.authService.CheckPermission(tokenStr, requiredRole)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "权限校验失败", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"has_permission": hasPermission})
}
