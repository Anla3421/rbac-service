package domain

import "context"

// UserRepository 用戶倉儲介面
type UserRepository interface {
	// GetByID 根據用戶ID獲取用戶信息
	GetByID(ctx context.Context, id string) (*User, error)
}
