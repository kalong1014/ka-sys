package domain

import "time"

// 用户模型
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // 注意：密码不返回给前端，用`-`标记
	Email     string    `json:"email"`
	Role      string    `json:"role"` // "user"（普通用户）或 "merchant"（商户用户）
	CreatedAt time.Time `json:"created_at"`
}

// 注册请求参数
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" validate:"required,oneof=user merchant"` // 注册时指定角色
}

// 注册响应结果
type RegisterResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// 登录请求参数
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// 登录响应结果（含JWT令牌）
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
