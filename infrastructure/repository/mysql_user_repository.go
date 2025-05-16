package repository

import (
	"context"
	"errors"
	"reflect"
	"strings"

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

// GetByUsername 根據用戶username獲取用戶信息
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

// UpdateUser 根據 username 更新用戶信息，僅更新非零值欄位
func (r *MySQLUserRepository) UpdateUser(ctx context.Context, username string, user *domain.User) error {
	// 使用 map 來存儲需要更新的欄位
	updateFields := make(map[string]interface{})

	// 使用反射檢查非零值欄位
	val := reflect.ValueOf(user).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Tag.Get("gorm")

		// 跳過空值和 ID 欄位
		if field.Interface() == reflect.Zero(field.Type()).Interface() ||
			fieldName == "primaryKey" ||
			fieldName == "-" {
			continue
		}

		// 根據 JSON tag 獲取欄位名
		jsonTag := typ.Field(i).Tag.Get("json")
		jsonName := strings.Split(jsonTag, ",")[0]

		// 將非零值欄位加入更新映射
		updateFields[jsonName] = field.Interface()
	}

	/* ------------------------------- 這邊會panic待修復 ------------------------------ */
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
	/* ------------------------------- 這邊會panic待修復 ------------------------------ */
	return nil
}
