package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"rbac-service/domain"
)

// AuthService 授權服務實作
type AuthService struct {
	authRepo domain.AuthRepository
}

// NewAuthService 創建新的 AuthService
func NewAuthService(authRepo domain.AuthRepository) *AuthService {
	return &AuthService{
		authRepo: authRepo,
	}
}

// GetUser 獲取用戶信息
func (s *AuthService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	// 驗證用戶ID
	id = strings.TrimSpace(id)
	if id == "" {
		return nil, domain.ErrInvalidUserID
	}

	// 調用倉儲層獲取用戶
	user, err := s.authRepo.GetByUsername(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login 處理使用者登入邏輯
func (s *AuthService) Login(username, password string) (string, error) {
	// 查詢使用者
	user, err := s.authRepo.GetByUsername(context.Background(), username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// 加密密碼並輸出 debug 資訊
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Debug 輸出
	if err != nil {
		return "", err
	}

	////// 測試用 print log，屆時要移除
	fmt.Println("輸入密碼加密結果:", string(hashedPassword))
	fmt.Println("資料庫密碼:", user.Password)

	// 驗證密碼
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// 產生 JWT token
	jwtKey := []byte("jwt_for_rcba_login")
	claims := jwt.MapClaims{
		"username": username,
		"role":     user.Roles,
		"exp":      jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("token generation failed")
	}

	// 使用 username 去更新剛剛建立的 jwt token
	needToupdate := map[string]interface{}{
		"Jwt": tokenString,
	}
	err = s.authRepo.UpdateUser(context.Background(), username, needToupdate)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	return tokenString, nil
}

// CheckPermission 檢查用戶是否有權限訪問特定資源
func (s *AuthService) CheckPermission(ctx context.Context, userID string, resource string, action string) (bool, error) {
	// 檢查輸入參數
	if userID == "" || resource == "" || action == "" {
		return false, errors.New("invalid input parameters")
	}

	// TODO: 這裡需要實現具體的權限檢查邏輯
	// 1. 獲取用戶的角色
	// 2. 獲取角色的權限列表
	// 3. 檢查是否包含所需的權限

	// 暫時返回 true 作為示例
	return true, nil
}
