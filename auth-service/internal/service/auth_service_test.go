package service

import (
	"auth-service/internal/domain"
	"context"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_Login(t *testing.T) {
	// 初始化测试服务
	secretKey := "test_secret_key"
	service := NewAuthService(secretKey, 24)

	// 注册测试用户
	user := domain.User{
		Username: "testuser",
		Password: "testpassword",
		Role:     domain.RoleUser,
	}
	err := service.RegisterUser(context.Background(), &user)
	assert.NoError(t, err)

	// 测试登录
	loginReq := domain.LoginRequest{
		Username: "testuser",
		Password: "testpassword",
	}

	tokenResp, err := service.Login(context.Background(), &loginReq)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenResp.Token)

	// 验证令牌
	claims := &domain.TokenClaims{}
	_, err = jwt.ParseWithClaims(tokenResp.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	assert.NoError(t, err)
	assert.Equal(t, user.Username, claims.UserID)
	assert.Equal(t, user.Role, claims.Role)
}

func TestAuthService_ValidateToken(t *testing.T) {
	secretKey := "test_secret_key"
	service := NewAuthService(secretKey, 24)

	// 生成有效令牌
	claims := &domain.TokenClaims{
		UserID: "testuser",
		Role:   domain.RoleUser,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secretKey))
	assert.NoError(t, err)

	// 验证有效令牌
	validatedClaims, err := service.ValidateToken(tokenStr)
	assert.NoError(t, err)
	assert.Equal(t, claims.UserID, validatedClaims.UserID)

	// 验证无效令牌
	_, err = service.ValidateToken("invalid_token")
	assert.Error(t, err)
}
