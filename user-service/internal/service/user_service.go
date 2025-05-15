package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"time"
	"user-service/internal/domain"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	users map[string]domain.User
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]domain.User),
	}
}

// 注册用户
func (s *UserService) Register(ctx context.Context, req *domain.RegisterRequest) (*domain.RegisterResponse, error) {
	log.Println("开始处理用户注册请求")
	// 校验参数（后续需添加验证中间件，这里先简单判断）
	if req.Username == "" || req.Password == "" || req.Email == "" {
		log.Println("注册参数缺失")
		return nil, errors.New("参数缺失")
	}

	// 生成用户ID
	userID := uuid.New().String()

	// 密码哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("密码哈希处理失败", err)
		return nil, err
	}

	// 创建用户对象
	newUser := domain.User{
		ID:        userID,
		Username:  req.Username,
		Password:  string(hashedPassword),
		Email:     req.Email,
		Role:      req.Role,
		CreatedAt: time.Now(),
	}

	// 存储到内存（使用用户名作为键）
	s.users[newUser.Username] = newUser
	log.Println("用户注册成功")

	// 返回注册结果（不包含密码）
	return &domain.RegisterResponse{
		ID:        userID,
		Username:  req.Username,
		Email:     req.Email,
		Role:      req.Role,
		CreatedAt: time.Now(),
	}, nil
}

// 登录验证
func (s *UserService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.LoginResponse, error) {
	log.Println("开始处理用户登录请求")
	// 通过用户名查找用户（修复：使用用户名作为键）
	user, exists := s.users[req.Username]
	if !exists {
		log.Println("用户不存在")
		return nil, errors.New("用户不存在")
	}

	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Println("密码错误")
		return nil, errors.New("密码错误")
	}

	// 生成JWT令牌（这里用随机字符串模拟，后续需实现真实JWT逻辑）
	token := generateMockToken(user.ID)
	log.Println("用户登录成功")

	return &domain.LoginResponse{
		Token: token,
		User: domain.User{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		},
	}, nil
}

// 模拟生成Token（生产环境需替换为JWT实现）
func generateMockToken(userID string) string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
