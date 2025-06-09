package delivery

import (
	"net/http"
	"rbac-service/domain"
	"rbac-service/usecase"
	"strings"

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

// AuthorizeRequest 授權請求參數
type AuthorizeRequest struct {
	Resource string `json:"resource" binding:"required"` // 要訪問的資源
	Action   string `json:"action" binding:"required"`   // 要執行的操作 (例如: read, write, delete)
}

// Login 處理用戶登錄請求
// @Summary 用戶登錄
// @Description 處理用戶登錄並返回訪問令牌
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "登錄請求參數"
// @Success 200 {object} map[string]interface{} "登錄成功"
// @Failure 400 {object} map[string]interface{} "無效的輸入或登錄失敗"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.NewErrorResponse("login failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})
}

// Logout 處理用戶登出請求
// @Summary 用戶登出
// @Description 處理用戶登出
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} map[string]interface{} "登出成功"
// @Failure 400 {object} map[string]interface{} "無效的輸入或登出失敗"
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	err := h.authService.Logout(c, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "logout failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

// Authorize 處理權限驗證的請求
// @Summary 驗證權限
// @Description 驗證用戶是否有權限訪問特定資源
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param request body AuthorizeRequest true "授權請求參數"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "驗證成功"
// @Failure 400 {object} map[string]interface{} "無效的請求參數"
// @Failure 401 {object} map[string]interface{} "未授權訪問"
// @Failure 403 {object} map[string]interface{} "權限不足"
// @Router /auth/authorize [post]
func (h *AuthHandler) Authorize(c *gin.Context) {
	var req AuthorizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewErrorResponse("Request Failed", "Invalid request parameters"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", "Please login first"))
		return
	}

	////// 待確認，是否要這樣用
	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", "Please login first"))
		return
	}

	hasPermission, err := h.authService.CheckPermission(c, userID.(string), token.(string), req.Resource, req.Action)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.NewErrorResponse("Permission Check Failed", err.Error()))
		return
	}

	if !hasPermission {
		c.JSON(http.StatusForbidden, domain.NewErrorResponse("Permission Denied", "No access to this resource"))
		return
	}
	authResponse := domain.AuthorizeResponse{
		Authorized: true,
		// TODO: Get expiration time from token
		ExpiresIn: 3600, // Example: 1 hour
	}

	c.JSON(http.StatusOK, domain.NewResponse("Authorization successful", authResponse))
}

// Refresh 處理刷新令牌的請求
// @Summary 刷新訪問令牌
// @Description 使用刷新令牌獲取新的訪問令牌
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "令牌刷新成功"
// @Failure 401 {object} map[string]interface{} "無效的刷新令牌"
// @Router /auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// Revoke 處理取消授權jwt的請求
// @Summary 撤銷訪問令牌
// @Description 撤銷指定的訪問令牌
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "令牌撤銷成功"
// @Failure 401 {object} map[string]interface{} "無效的令牌"
// @Router /auth/revoke [post]
func (h *AuthHandler) Revoke(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// BatchRevoke 處理批量取消授權jwt的請求
// @Summary 批量撤銷訪問令牌
// @Description 批量撤銷多個訪問令牌
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "批量撤銷成功"
// @Failure 401 {object} map[string]interface{} "無效的令牌"
// @Router /auth/batchRevoke [post]
func (h *AuthHandler) BatchRevoke(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
