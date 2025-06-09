package domain

import "context"

// BaseRepository 定義基礎倉儲方法
type BaseRepository interface {
	GetByID(ctx context.Context, id string) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	UpdateUser(ctx context.Context, username string, updateFields map[string]interface{}) error
	DeleteUserJwt(ctx context.Context, jwt string) error
	CreateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, username string) error
}

type UserRepository interface {
	BaseRepository
	// 可以添加特定於用戶的額外方法
}

type AuthRepository interface {
	BaseRepository
}
