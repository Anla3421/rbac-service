package middleware

import (
	"net/http"
	"strings"

	"rbac-service/domain"
	"rbac-service/infrastructure/utils"
	"rbac-service/usecase"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware 創建 JWT 驗證中間件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 從 Header 提取 token
		token := c.GetHeader("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")

		// 2. 檢查 token 是否為空
		if token == "" {
			c.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", "Missing token"))
			c.Abort()
			return
		}

		// 3. 檢查 token 是否過期
		if utils.IsTokenExpired(token) {
			c.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", "Token expired"))
			c.Abort()
			return
		}

		// 4. 解析 token
		claims, err := utils.ParseJWTToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", "Invalid token"))
			c.Abort()
			return
		}

		// 5. 提取用戶名
		username, ok := claims["username"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", "Invalid token claims"))
			c.Abort()
			return
		}

		// 6. 檢查 token 是否與數據庫一致
		isValid, err := utils.CompareJWTToken(c, username, token)
		if err != nil || !isValid {
			c.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", "Token invalidated"))
			c.Abort()
			return
		}

		// 7. 將用戶信息存入 context
		c.Set("userID", username)
		c.Set("token", token)

		// 8. 繼續處理請求
		c.Next()
	}
}

// PermissionMiddleware 權限中間件
func PermissionMiddleware(authService *usecase.AuthService, resource, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 從 context 獲取用戶信息
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", "Please login first"))
			c.Abort()
			return
		}

		token, exists := c.Get("token")
		if !exists {
			c.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", "Please login first"))
			c.Abort()
			return
		}

		// 檢查權限
		hasPermission, err := authService.CheckPermission(c, userID.(string), token.(string), resource, action)
		if err != nil {
			c.JSON(http.StatusForbidden, domain.NewErrorResponse("Permission Denied", err.Error()))
			c.Abort()
			return
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, domain.NewErrorResponse("Permission Denied", "No access to this resource"))
			c.Abort()
			return
		}

		c.Next()
	}
}
