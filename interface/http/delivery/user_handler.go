package delivery

import (
	"context"
	"net/http"
	"rbac-service/domain"
	"rbac-service/usecase"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 定義更新請求的結構體，根據目前定義的 schema 只有密碼能修改，若後續 schema 有變更再一起修改
type UpdateUserRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// UserHandler 處理用戶相關的 HTTP 請求
type UserHandler struct {
	userService *usecase.UserService
}

// NewUserHandler 創建新的 UserHandler
func NewUserHandler(userService *usecase.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Create 處理創建用戶的請求
// @Summary 創建新用戶
// @Description 使用提供的用戶名和密碼創建新用戶
// @Tags Users
// @Accept json
// @Produce json
// @Param user body LoginRequest true "用戶創建信息"
// @Success 201 {object} map[string]interface{} "用戶創建成功"
// @Failure 400 {object} map[string]string "參數驗證失敗"
// @Failure 409 {object} map[string]string "用戶名已存在"
// @Failure 500 {object} map[string]string "服務器內部錯誤"
// @Router /users/registry [post]
func (h *UserHandler) Create(c *gin.Context) {
	////// 定義用戶創建請求的結構體，要 loging 共用還是分開？
	// type CreateUserRequest struct {
	// 	Username string `json:"username" binding:"required" example:"johndoe"`
	// 	Password string `json:"password" binding:"required,min=5" example:"password123"`
	// }

	// 解析請求數據
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 統一返回400，避免被猜測
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "參數驗證失敗",
		})
		return
	}

	// 檢查用戶名是否已存在
	existingUser, _ := h.userService.GetByUsername(context.Background(), req.Username)
	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "用戶名已存在",
		})
		return
	}

	// 密碼加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "密碼處理失敗",
		})
		return
	}

	// 創建新用戶
	newUser := &domain.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	// 調用用戶服務創建用戶
	createdUser, err := h.userService.CreateUser(context.Background(), newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "創建用戶失敗",
		})
		return
	}

	// 返回創建成功的用戶信息
	c.JSON(http.StatusCreated, gin.H{
		"user": gin.H{
			"id":       createdUser.ID,
			"username": createdUser.Username,
		},
	})
}

// List 處理列出用戶的請求
func (h *UserHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// Get 處理獲取單個用戶的請求
// @Summary 獲取用戶詳情
// @Description 根據用戶ID獲取用戶詳細信息
// @Tags Users
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "用戶ID"
// @Success 200 {object} domain.User "成功獲取用戶信息"
// @Failure 400 {object} map[string]string "無效的用戶ID"
// @Failure 404 {object} map[string]string "用戶未找到"
// @Failure 500 {object} map[string]string "服務器內部錯誤"
// @Router /users/{id} [get]
func (h *UserHandler) Get(c *gin.Context) {
	id := c.Param("id")

	// 調用服務層獲取用戶
	user, err := h.userService.GetUser(c, id)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "用戶未找到",
			})
			return
		case domain.ErrInvalidUserID:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "無效的用戶ID",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "伺服器錯誤",
			})
			return
		}
	}

	// 返回用戶信息
	c.JSON(http.StatusOK, user)
}

// Update 處理更新用戶的請求
// @Summary 更新用戶信息
// @Description 根據用戶名更新用戶密碼或其他信息，根據目前定義的 schema 只有密碼能修改，若後續 schema 有變更再一起修改
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "用戶ID"
// @Param user body UpdateUserRequest true "用戶更新信息"
// @Success 200 {object} map[string]interface{} "成功更新用戶信息"
// @Failure 400 {object} map[string]string "參數驗證失敗或無效的用戶ID"
// @Failure 404 {object} map[string]string "用戶未找到"
// @Failure 500 {object} map[string]string "服務器內部錯誤"
// @Router /users [put]
func (h *UserHandler) Update(c *gin.Context) {

	// 解析請求數據
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "參數驗證失敗",
		})
		return
	}

	// 準備更新的用戶資訊
	updateUser := &domain.User{
		Username: req.Username,
		Password: req.Password,
	}

	// 如果提供了新的密碼，則加密
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "密碼處理失敗",
			})
			return
		}
		updateUser.Password = string(hashedPassword)
	}

	// 調用用戶服務更新用戶
	updatedUser, err := h.userService.UpdateUser(c, updateUser)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "用戶未找到",
			})
			return
		case domain.ErrInvalidUserID:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "無效的用戶ID",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "伺服器錯誤",
			})
			return
		}
	}

	// 返回更新成功的用戶信息
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":       updatedUser.ID,
			"username": updatedUser.Username,
		},
	})
}

// Delete 處理刪除用戶的請求
// @Summary 刪除用戶
// @Description 根據用戶名直接刪除用戶
// @Tags Users
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} map[string]string "用戶成功刪除"
// @Failure 404 {object} map[string]string "用戶未找到"
// @Failure 500 {object} map[string]string "服務器內部錯誤"
// @Router /users/ [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	// 準備更新的用戶資訊
	updateUser := &domain.User{
		Username: c.GetString("username"),
		Jwt:      c.GetString("token"),
	}
	err := h.userService.DeleteUser(c, updateUser)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": domain.ErrUserNotFound.Error(),
			})
			return
		case domain.ErrInvalidJwt:
			c.JSON(http.StatusNotFound, gin.H{
				"error": domain.ErrInvalidJwt.Error(),
			})
			return

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": domain.ErrInternalServerError.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "用戶刪除成功",
	})
}
