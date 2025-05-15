package domain

import (
	"time"
)

// 路由配置
type Route struct {
	ID           string    `json:"id"`
	Path         string    `json:"path"`          // 匹配路径
	Method       string    `json:"method"`        // HTTP方法
	ServiceName  string    `json:"service_name"`  // 目标服务名称
	ServiceAddr  string    `json:"service_addr"`  // 目标服务地址
	Timeout      int       `json:"timeout"`       // 请求超时时间（毫秒）
	RetryTimes   int       `json:"retry_times"`   // 重试次数
	RateLimit    int       `json:"rate_limit"`    // 限流（请求/秒）
	AuthRequired bool      `json:"auth_required"` // 是否需要认证
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// 认证配置
type AuthConfig struct {
	ID          string    `json:"id"`
	ServiceName string    `json:"service_name"` // 认证服务名称
	ServiceAddr string    `json:"service_addr"` // 认证服务地址
	TokenHeader string    `json:"token_header"` // 令牌头名称
	PublicPaths []string  `json:"public_paths"` // 公共路径（无需认证）
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// 创建路由请求
type CreateRouteRequest struct {
	Path         string `json:"path" validate:"required"`
	Method       string `json:"method" validate:"required"`
	ServiceName  string `json:"service_name" validate:"required"`
	ServiceAddr  string `json:"service_addr" validate:"required"`
	Timeout      int    `json:"timeout"`
	RetryTimes   int    `json:"retry_times"`
	RateLimit    int    `json:"rate_limit"`
	AuthRequired bool   `json:"auth_required"`
}

// 创建认证配置请求
type CreateAuthConfigRequest struct {
	ServiceName string   `json:"service_name" validate:"required"`
	ServiceAddr string   `json:"service_addr" validate:"required"`
	TokenHeader string   `json:"token_header"`
	PublicPaths []string `json:"public_paths"`
}
