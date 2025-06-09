package usecase

import (
	"context"
	"errors"
	"strings"

	"rbac-service/domain"
	"rbac-service/infrastructure/utils"
)

// UserService 用戶服務實作
type UserService struct {
	repo domain.UserRepository
}

// NewUserService 創建用戶服務
func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUser 獲取用戶信息
func (s *UserService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	// 驗證用戶ID
	id = strings.TrimSpace(id)
	if id == "" {
		return nil, domain.ErrInvalidUserID
	}

	// 調用倉儲層獲取用戶
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetByUsername 獲取用戶信息
func (s *UserService) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	// 驗證用戶ID
	username = strings.TrimSpace(username)
	if username == "" {
		return nil, domain.ErrInvalidUserID
	}

	// 調用倉儲層獲取用戶
	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser 創建用戶
func (s *UserService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	// 調用倉儲層創建用戶
	createdUser, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

// UpdateUser 更新用戶資料
func (s *UserService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	// 檢查用戶名是否為空
	if user.Username == "" {
		return nil, domain.ErrInvalidUserID
	}

	// 準備更新的欄位
	updateFields := make(map[string]interface{})

	// 如果有密碼，加入更新欄位
	if user.Password != "" {
		updateFields["password"] = user.Password
	}

	// 如果沒有可更新的欄位，直接返回
	if len(updateFields) == 0 {
		return nil, errors.New("沒有可更新的欄位")
	}

	// 調用倉儲層更新用戶
	err := s.repo.UpdateUser(ctx, user.Username, updateFields)
	if err != nil {
		return nil, err
	}

	// 重新獲取更新後的用戶信息
	updatedUser, err := s.repo.GetByUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// DeleteUser 刪除用戶
func (s *UserService) DeleteUser(ctx context.Context, user *domain.User) error {
	// 檢查 jwt 是否跟 db 內相同，final check
	isTrue, err := utils.CompareJWTToken(ctx, user.Username, user.Jwt)
	if err != nil || !isTrue {
		return domain.ErrInvalidJwt
	}
	// 調用倉儲層刪除用戶
	err = s.repo.DeleteUser(ctx, user.Username)
	if err != nil {
		return err
	}

	return nil
}
