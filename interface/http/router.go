package http

import (
	"net/http"
	_ "rbac-service/docs"
	"rbac-service/interface/http/delivery"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter 設置路由
func SetupRouter(
	r *gin.Engine,
	userHandler *delivery.UserHandler,
	authHandler *delivery.AuthHandler,
) {
	// Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 設定基本路由群組
	v1 := r.Group("/v1")
	{
		// 用戶管理路由
		userGroup := v1.Group("/users")
		{
			userGroup.POST("/create", userHandler.Create)

			// @Summary 列出用戶
			// @Description 獲取用戶列表
			// @Tags Users
			// @Produce json
			// @Success 200 {object} map[string]string
			// @Router /users [get]
			userGroup.GET("", userHandler.List)

			userGroup.GET("/:id", userHandler.Get)

			// @Summary 更新用戶
			// @Description 更新用戶信息
			// @Tags Users
			// @Accept json
			// @Produce json
			// @Param id path string true "用戶ID"
			// @Success 200 {object} map[string]string
			// @Router /users/{id} [put]
			userGroup.PUT("/:id", userHandler.Update)

			// @Summary 刪除用戶
			// @Description 根據ID刪除用戶
			// @Tags Users
			// @Produce json
			// @Param id path string true "用戶ID"
			// @Success 200 {object} map[string]string
			// @Router /users/{id} [delete]
			userGroup.DELETE("/:id", userHandler.Delete)
		}

		// 角色管理路由
		roleGroup := v1.Group("/roles")
		{
			// @Summary 創建角色
			// @Description 創建新的角色
			// @Tags Roles
			// @Accept json
			// @Produce json
			// @Success 200 {object} map[string]string
			// @Router /roles [post]
			roleGroup.POST("", createRole)

			// @Summary 列出角色
			// @Description 獲取角色列表
			// @Tags Roles
			// @Produce json
			// @Success 200 {object} map[string]string
			// @Router /roles [get]
			roleGroup.GET("", listRoles)

			// @Summary 獲取角色
			// @Description 根據ID獲取角色詳情
			// @Tags Roles
			// @Produce json
			// @Param id path string true "角色ID"
			// @Success 200 {object} map[string]string
			// @Router /roles/{id} [get]
			roleGroup.GET("/:id", getRole)

			// @Summary 更新角色
			// @Description 更新角色信息
			// @Tags Roles
			// @Accept json
			// @Produce json
			// @Param id path string true "角色ID"
			// @Success 200 {object} map[string]string
			// @Router /roles/{id} [put]
			roleGroup.PUT("/:id", updateRole)

			// @Summary 刪除角色
			// @Description 根據ID刪除角色
			// @Tags Roles
			// @Produce json
			// @Param id path string true "角色ID"
			// @Success 200 {object} map[string]string
			// @Router /roles/{id} [delete]
			roleGroup.DELETE("/:id", deleteRole)
		}

		// 權限管理路由
		permissionGroup := v1.Group("/permissions")
		{
			// @Summary 創建權限
			// @Description 創建新的權限
			// @Tags Permissions
			// @Accept json
			// @Produce json
			// @Success 200 {object} map[string]string
			// @Router /permissions [post]
			permissionGroup.POST("", createPermission)

			// @Summary 列出權限
			// @Description 獲取權限列表
			// @Tags Permissions
			// @Produce json
			// @Success 200 {object} map[string]string
			// @Router /permissions [get]
			permissionGroup.GET("", listPermissions)

			// @Summary 獲取權限
			// @Description 根據ID獲取權限詳情
			// @Tags Permissions
			// @Produce json
			// @Param id path string true "權限ID"
			// @Success 200 {object} map[string]string
			// @Router /permissions/{id} [get]
			permissionGroup.GET("/:id", getPermission)

			// @Summary 更新權限
			// @Description 更新權限信息
			// @Tags Permissions
			// @Accept json
			// @Produce json
			// @Param id path string true "權限ID"
			// @Success 200 {object} map[string]string
			// @Router /permissions/{id} [put]
			permissionGroup.PUT("/:id", updatePermission)

			// @Summary 刪除權限
			// @Description 根據ID刪除權限
			// @Tags Permissions
			// @Produce json
			// @Param id path string true "權限ID"
			// @Success 200 {object} map[string]string
			// @Router /permissions/{id} [delete]
			permissionGroup.DELETE("/:id", deletePermission)
		}

		// 授權管理路由
		authGroup := v1.Group("/auth")
		{
			// 登入
			authGroup.POST("login", authHandler.Login)
			// 登出
			authGroup.POST("logout", authHandler.Logout)
			// 權限驗證
			authGroup.POST("authorize", authHandler.Authorize)
			// 刷新令牌
			authGroup.POST("refresh", authHandler.Refresh)
			// 取消授權jwt
			authGroup.POST("revoke", authHandler.Revoke)
			// 批量取消授權jwt
			authGroup.POST("batch-revoke", authHandler.BatchRevoke)
		}
	}

	// 健康檢查路由
	// @Summary 健康檢查
	// @Description 檢查服務是否正常運行
	// @Tags Health
	// @Produce json
	// @Success 200 {object} map[string]string
	// @Router /health [get]
	r.GET("/health", healthCheck)
}

// 預設處理函式，暫時返回 ok
func createUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func listUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func updateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func deleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func createRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func listRoles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func getRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func updateRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func deleteRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func createPermission(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func listPermissions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func getPermission(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func updatePermission(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func deletePermission(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
