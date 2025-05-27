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

// GetByID 根據用戶 ID 獲取用戶信息
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

// GetByUsername 根據用戶 username 獲取用戶信息
func (r *MySQLUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	result := r.db.WithContext(ctx).Where("username = ?", username).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}

// UpdateUser 根據 username 更新用戶信息，可 partial update
func (r *MySQLUserRepository) UpdateUser(ctx context.Context, username string, updateFields map[string]interface{}) error {
	// 執行更新
	result := r.db.WithContext(ctx).
		Model(&domain.User{}).
		Where("username = ?", username).
		Updates(updateFields)

	if result.Error != nil {
		return result.Error
	}

	// 檢查是否有實際更新
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

// CreateUser 創建用戶
func (r *MySQLUserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	// 只創建指定的欄位，排除 Roles
	result := r.db.WithContext(ctx).Omit("Roles").Create(user)
	return user, result.Error
}
