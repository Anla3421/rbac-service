package repository

import (
	"context"
	"errors"

	"rbac-service/domain"

	"gorm.io/gorm"
)

// MySQLUserRepository MySQL 用戶倉儲實作
type MySQLUserRepository struct {
	db *gorm.DB
}

// NewMySQLUserRepository 創建 MySQL 用戶倉儲
func NewMySQLUserRepository(db *gorm.DB) domain.UserRepository {
	return &MySQLUserRepository{db: db}
}

// GetByID 根據用戶ID獲取用戶信息
func (r *MySQLUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}
