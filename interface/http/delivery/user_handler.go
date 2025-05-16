package delivery

import (
	"net/http"
	"rbac-service/domain"
	"rbac-service/usecase"

	"github.com/gin-gonic/gin"
)

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
func (h *UserHandler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
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
func (h *UserHandler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// Delete 處理刪除用戶的請求
func (h *UserHandler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
