package domain

import "errors"

var (
	// ErrUserNotFound 用戶未找到錯誤
	ErrUserNotFound = errors.New("user not found")

	// ErrInvalidUserID 無效的用戶ID
	ErrInvalidUserID = errors.New("invalid user ID")
)
