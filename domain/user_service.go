package domain

import "context"

// UserService 用戶服務介面
type UserService interface {
	// GetUser 獲取用戶信息
	GetUser(ctx context.Context, id string) (*User, error)
}
