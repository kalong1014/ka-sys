package repository

import (
	"auth-service/internal/domain"
	"context"
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserMySQLRepo struct {
	db *gorm.DB
}

func NewUserMySQLRepo(dsn string) (*UserMySQLRepo, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		return nil, err
	}

	return &UserMySQLRepo{db: db}, nil
}

// 创建用户
func (r *UserMySQLRepo) CreateUser(ctx context.Context, user *domain.User) error {
	return r.db.Create(user).Error
}

// 获取用户
func (r *UserMySQLRepo) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "username = ?", username).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}
