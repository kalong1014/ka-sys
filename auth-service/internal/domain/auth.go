package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 用户角色
const (
	RoleAdmin = "admin" // 管理员
	RoleUser  = "user"  // 普通用户
	RoleGuest = "guest" // 访客
)

// 用户信息
type User struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=6,max=64"` // 存储哈希值
	Email    string `json:"email" validate:"email"`
	Role     string `json:"role" validate:"required,oneof=admin user guest"`
}

// 登录请求
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// JWT令牌Claims
type TokenClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// 令牌响应
type TokenResponse struct {
	Token        string    `json:"token"`
	ExpireAt     time.Time `json:"expire_at"`
	RefreshToken string    `json:"refresh_token"`
}
