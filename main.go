package main

import (
	"log"
	_ "rbac-service/docs"
	"rbac-service/interface/http"

	"rbac-service/infrastructure/database"
	"rbac-service/infrastructure/repository"
	"rbac-service/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// @title           RBAC Service API
	// @version         1.0
	// @description     RBAC 權限管理服務 API 文檔
	// @host            localhost:8080
	// @BasePath        /v1

	// 初始化資料庫
	dbManager := database.NewDatabaseManager()

	// 載入資料庫配置 - 這是關鍵步驟
	if err := dbManager.LoadConfigs(); err != nil {
		log.Fatalf("Failed to load database configuration: %v", err)
	}

	// 獲取主資料庫連接
	primaryDB, err := dbManager.GetDatabase("rbac")
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}

	// 創建依賴鏈
	userRepo := repository.NewMySQLUserRepository(primaryDB)
	userService := usecase.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)

	// 設置路由
	r := gin.Default()
	r.Use(cors.Default())
	http.SetupRouter(r, userHandler)

	// 啟動伺服器
	///// "localhost:5001" 測試用，避免每次都要按防火牆擋案允許，應用":5001"
	if err := r.Run(":5001"); err != nil {
		log.Fatalf("Server startup failed: %v", err)
	}
}
