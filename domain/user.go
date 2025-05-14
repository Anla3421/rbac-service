package domain

import "time"

// User 領域模型
type User struct {
	UID       string    `json:"uid"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository 用戶倉儲介面
type UserRepository interface {
	Create(user *User) error
	GetByID(uid string) (*User, error)
	Update(user *User) error
	Delete(uid string) error
	List(page, pageSize int) ([]*User, int, error)
}

// UserService 用戶服務介面
type UserService interface {
	CreateUser(user *User) error
	GetUser(uid string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(uid string) error
	ListUsers(page, pageSize int) ([]*User, int, error)
}
