package main

import (
	"log"
	_ "rbac-service/docs"
	httpHandler "rbac-service/interface/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           RBAC Service API
// @version         1.0
// @description     RBAC 權限管理服務 API 文檔
// @host            localhost:8080
// @BasePath        /v1
func main() {
	r := gin.Default()

	// 初始化路由
	httpHandler.SetupRouter(r)

	// Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 啟動伺服器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server startup failed: %v", err)
	}
}
