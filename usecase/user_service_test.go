package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"rbac-service/domain"
)

// MockUserRepository 模擬 UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

// 實現 GetByUsername 方法以滿足接口要求
func (m *MockUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, username string, updates map[string]interface{}) error {
	args := m.Called(ctx, username, updates)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, username string) error {
	args := m.Called(ctx, username)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUserJwt(ctx context.Context, jwt string) error {
	args := m.Called(ctx, jwt)
	return args.Error(0)
}

func TestUserService_GetUser_SuccessfulRetrieval(t *testing.T) {
	// 準備測試數據
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)

	userID := "testuser123"
	expectedUser := &domain.User{
		Username: userID,
		Roles:    []string{"user"},
	}

	// 設定模擬行為
	mockRepo.On("GetByID", mock.Anything, userID).Return(expectedUser, nil)

	// 執行獲取用戶
	user, err := userService.GetUser(context.Background(), userID)

	// 斷言
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUser_EmptyUserID(t *testing.T) {
	// 準備測試數據
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)

	// 執行獲取用戶
	user, err := userService.GetUser(context.Background(), "")

	// 斷言
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, domain.ErrInvalidUserID, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUser_UserNotFound(t *testing.T) {
	// 準備測試數據
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)

	userID := "nonexistentuser"
	expectedError := errors.New("user not found")

	// 設定模擬行為
	mockRepo.On("GetByID", mock.Anything, userID).Return(nil, expectedError)

	// 執行獲取用戶
	user, err := userService.GetUser(context.Background(), userID)

	// 斷言
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetByUsername_SuccessfulRetrieval(t *testing.T) {
	// 準備測試數據
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)

	username := "testuser123"
	expectedUser := &domain.User{
		Username: username,
		Roles:    []string{"user"},
	}

	// 設定模擬行為
	mockRepo.On("GetByUsername", mock.Anything, username).Return(expectedUser, nil)

	// 執行獲取用戶
	user, err := userService.GetByUsername(context.Background(), username)

	// 斷言
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetByUsername_EmptyUsername(t *testing.T) {
	// 準備測試數據
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)

	// 執行獲取用戶
	user, err := userService.GetByUsername(context.Background(), "")

	// 斷言
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, domain.ErrInvalidUserID, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetByUsername_UserNotFound(t *testing.T) {
	// 準備測試數據
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)

	username := "nonexistentuser"
	expectedError := errors.New("user not found")

	// 設定模擬行為
	mockRepo.On("GetByUsername", mock.Anything, username).Return(nil, expectedError)

	// 執行獲取用戶
	user, err := userService.GetByUsername(context.Background(), username)

	// 斷言
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
