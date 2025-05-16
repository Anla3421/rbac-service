package usecase

import (
	"context"
	"strings"

	"rbac-service/domain"
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
