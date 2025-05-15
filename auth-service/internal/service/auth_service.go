package service

import (
	"auth-service/internal/domain"
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	secretKey     string                 // JWT密钥
	tokenLifetime time.Duration          // 令牌有效期
	users         map[string]domain.User // 用户存储（内存模拟）
}

func NewAuthService(secretKey string, tokenLifetime int) *AuthService {
	return &AuthService{
		secretKey:     secretKey,
		tokenLifetime: time.Duration(tokenLifetime) * time.Hour,
		users:         make(map[string]domain.User),
	}
}

// 用户注册
func (s *AuthService) RegisterUser(ctx context.Context, user *domain.User) error {
	log.Printf("用户注册: 用户名=%s", user.Username)

	// 检查用户名是否已存在
	if _, exists := s.users[user.Username]; exists {
		return errors.New("用户名已存在")
	}

	// 哈希密码
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 生成用户ID
	user.ID = generateUserID(user.Username)

	// 保存用户（内存存储，实际应使用数据库）
	user.Password = string(hashedPwd) // 存储哈希值
	s.users[user.Username] = *user

	log.Printf("用户注册成功: ID=%s", user.ID)
	return nil
}

// 用户登录
func (s *AuthService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.TokenResponse, error) {
	log.Printf("用户登录: 用户名=%s", req.Username)

	// 查找用户
	user, exists := s.users[req.Username]
	if !exists {
		return nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成JWT令牌
	return s.generateTokens(user.ID, user.Role)
}

// 生成令牌
func (s *AuthService) generateTokens(userID, role string) (*domain.TokenResponse, error) {
	expireAt := time.Now().Add(s.tokenLifetime)

	// 访问令牌Claims
	claims := &domain.TokenClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// 生成JWT令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return nil, err
	}

	// 简化处理：刷新令牌暂用相同逻辑（实际应独立生成）
	refreshTokenStr, err := token.SignedString([]byte(s.secretKey + "refresh"))
	if err != nil {
		return nil, err
	}

	return &domain.TokenResponse{
		Token:        tokenStr,
		ExpireAt:     expireAt,
		RefreshToken: refreshTokenStr,
	}, nil
}

// 验证令牌
func (s *AuthService) ValidateToken(tokenStr string) (*domain.TokenClaims, error) {
	claims := &domain.TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("无效令牌")
	}

	return claims, nil
}

// 生成用户ID（简化实现）
func generateUserID(username string) string {
	// 实际应使用UUID生成
	return "usr_" + username[:3]
}

// 检查用户权限
func (s *AuthService) CheckPermission(tokenStr, requiredRole string) (bool, error) {
	claims, err := s.ValidateToken(tokenStr)
	if err != nil {
		return false, err
	}

	// 检查角色权限
	if requiredRole == domain.RoleAdmin {
		return claims.Role == domain.RoleAdmin, nil
	}

	if requiredRole == domain.RoleUser {
		return claims.Role == domain.RoleUser || claims.Role == domain.RoleAdmin, nil
	}

	// 其他角色（如访客）
	return true, nil
}
