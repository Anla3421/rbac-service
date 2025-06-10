package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"rbac-service/domain"
	"rbac-service/infrastructure/utils"
)

// MockAuthRepository 模擬 AuthRepository
type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockAuthRepository) UpdateUser(ctx context.Context, username string, updates map[string]interface{}) error {
	args := m.Called(ctx, username, updates)
	return args.Error(0)
}

func (m *MockAuthRepository) DeleteUserJwt(ctx context.Context, jwt string) error {
	args := m.Called(ctx, jwt)
	return args.Error(0)
}

// 修正 CreateUser 方法的簽名
func (m *MockAuthRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockAuthRepository) DeleteUser(ctx context.Context, username string) error {
	args := m.Called(ctx, username)
	return args.Error(0)
}

func (m *MockAuthRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestLogin_SuccessfulLogin(t *testing.T) {
	// 準備測試數據
	mockRepo := new(MockAuthRepository)
	authService := NewAuthService(mockRepo)

	username := "testuser"
	rawPassword := "password123"
	hashedPassword, _ := utils.HashPassword(rawPassword)

	mockUser := &domain.User{
		Username: username,
		Password: hashedPassword,
		Roles:    []string{"user"},
	}

	// 設定模擬行為
	mockRepo.On("GetByUsername", mock.Anything, username).Return(mockUser, nil)
	mockRepo.On("UpdateUser", mock.Anything, username, mock.Anything).Return(nil)

	// 執行登入
	token, err := authService.Login(username, rawPassword)

	// 斷言
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLogin_InvalidUsername(t *testing.T) {
	// 準備測試數據
	mockRepo := new(MockAuthRepository)
	authService := NewAuthService(mockRepo)

	username := "nonexistentuser"

	// 設定模擬行為
	mockRepo.On("GetByUsername", mock.Anything, username).Return(nil, errors.New("user not found"))

	// 執行登入
	token, err := authService.Login(username, "anypassword")

	// 斷言
	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Contains(t, err.Error(), "invalid credentials")
	mockRepo.AssertExpectations(t)
}

func TestLogin_WrongPassword(t *testing.T) {
	// 準備測試數據
	mockRepo := new(MockAuthRepository)
	authService := NewAuthService(mockRepo)

	username := "testuser"
	correctPassword := "correctpassword"
	wrongPassword := "wrongpassword"
	hashedPassword, _ := utils.HashPassword(correctPassword)

	mockUser := &domain.User{
		Username: username,
		Password: hashedPassword,
		Roles:    []string{"user"},
	}

	// 設定模擬行為
	mockRepo.On("GetByUsername", mock.Anything, username).Return(mockUser, nil)

	// 執行登入
	token, err := authService.Login(username, wrongPassword)

	// 斷言
	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Contains(t, err.Error(), "invalid credentials")
	mockRepo.AssertExpectations(t)
}

func TestLogout_SuccessfulLogout(t *testing.T) {
	// 準備測試數據
	mockRepo := new(MockAuthRepository)
	authService := NewAuthService(mockRepo)

	jwt := "some-valid-jwt-token"

	// 設定模擬行為：預期 DeleteUserJwt 被呼叫並回傳 nil
	mockRepo.On("DeleteUserJwt", mock.Anything, jwt).Return(nil)

	// 執行登出
	err := authService.Logout(context.Background(), jwt)

	// 斷言
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLogout_FailedLogout(t *testing.T) {
	// 準備測試數據
	mockRepo := new(MockAuthRepository)
	authService := NewAuthService(mockRepo)

	jwt := "some-invalid-jwt-token"
	expectedErr := errors.New("failed to delete jwt")

	// 設定模擬行為：預期 DeleteUserJwt 被呼叫並回傳錯誤
	mockRepo.On("DeleteUserJwt", mock.Anything, jwt).Return(expectedErr)

	// 執行登出
	err := authService.Logout(context.Background(), jwt)

	// 斷言
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	mockRepo.AssertExpectations(t)
}
