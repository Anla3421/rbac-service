package delivery

import (
	"net/http"
	"rbac-service/domain"
	"rbac-service/usecase"

	"github.com/gin-gonic/gin"
)

// AuthHandler 處理用戶相關的 HTTP 請求
type AuthHandler struct {
	authService *usecase.AuthService
}

// NewAuthHandler 創建新的 AuthHandler
func NewAuthHandler(authService *usecase.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Authorize 處理權限驗證的請求
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})
}

// Authorize 處理權限驗證的請求
// @Router /auth/authorize [post]
func (h *AuthHandler) Authorize(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// Refresh 處理刷新令牌的請求
// @Router /auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// Revoke 處理取消授權jwt的請求
// @Router /auth/revoke [post]
func (h *AuthHandler) Revoke(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// BatchRevoke 處理批量取消授權jwt的請求
// @Router /auth/batchRevoke [post]
func (h *AuthHandler) BatchRevoke(c *gin.Context) {
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
func (h *AuthHandler) Get(c *gin.Context) {
	id := c.Param("id")

	// 調用服務層獲取用戶
	user, err := h.authService.GetUser(c, id)
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
