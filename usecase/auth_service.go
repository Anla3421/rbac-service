package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"rbac-service/domain"
	"rbac-service/infrastructure/utils"
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
	tokenString, err := utils.GenerateJWTToken(username, user.Roles)
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

// Logout 處理使用者登出邏輯
func (s *AuthService) Logout(ctx context.Context, jwt string) error {
	return s.authRepo.DeleteUserJwt(ctx, jwt)
}

// CheckPermission 檢查用戶是否有權限訪問特定資源
func (s *AuthService) CheckPermission(ctx context.Context, userID string, token string, resource string, action string) (bool, error) {
	// 1. 檢查輸入參數
	if userID == "" || token == "" || resource == "" || action == "" {
		return false, errors.New("invalid input parameters")
	}

	// 2. 先解析 token
	claims, err := utils.ParseJWTToken(token)
	if err != nil {
		return false, errors.New("invalid token")
	}
	fmt.Println("解析後的 token:", claims)

	// 3. 從資料庫取出用戶當前的 JWT
	user, err := s.authRepo.GetByUsername(ctx, userID)
	if err != nil {
		return false, errors.New("user not found")
	}

	// 4. 比對 token
	if user.Jwt != token {
		return false, errors.New("token has been invalidated")
	}

	// 5. 進行權限檢查的邏輯
	// TODO: 實現具體的權限檢查
	return true, nil
}
