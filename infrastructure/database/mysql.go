package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseConfig 單一資料庫配置
type DatabaseConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Database  string `json:"database"`
	Type      string `json:"type"`
	Charset   string `json:"charset,omitempty"`
	ParseTime bool   `json:"parseTime,omitempty"`
}

// MultiDatabaseConfig 多資料庫配置
type MultiDatabaseConfig struct {
	Databases map[string]DatabaseConfig `json:"databases"`
}

// DatabaseManager 管理多個資料庫連接
type DatabaseManager struct {
	configs map[string]*gorm.DB
	mu      sync.RWMutex
}

// NewDatabaseManager 創建資料庫管理器
func NewDatabaseManager() *DatabaseManager {
	return &DatabaseManager{
		configs: make(map[string]*gorm.DB),
	}
}

// LoadConfigs 載入所有資料庫配置
func (dm *DatabaseManager) LoadConfigs() error {
	data, err := os.ReadFile("configs/database.json")
	if err != nil {
		return fmt.Errorf("讀取配置檔案失敗: %v", err)
	}

	var multiConfig MultiDatabaseConfig
	if err := json.Unmarshal(data, &multiConfig); err != nil {
		return fmt.Errorf("解析配置失敗: %v", err)
	}

	for name, config := range multiConfig.Databases {
		db, err := dm.connectDatabase(name, config)
		if err != nil {
			log.Printf("連接 %s 資料庫失敗: %v", name, err)
			continue
		}
		dm.configs[name] = db
	}

	return nil
}

// connectDatabase 連接單一資料庫
func (dm *DatabaseManager) connectDatabase(name string, config DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.Charset,
		config.ParseTime,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("連接 %s 資料庫失敗: %v", name, err)
	}

	return db, nil
}

// GetDatabase 獲取指定名稱的資料庫連接
func (dm *DatabaseManager) GetDatabase(name string) (*gorm.DB, error) {
	dm.mu.RLock()
	defer dm.mu.RUnlock()

	// 檢查是否已載入配置
	if len(dm.configs) == 0 {
		return nil, fmt.Errorf("資料庫配置尚未載入，請先調用 LoadConfigs()")
	}

	db, exists := dm.configs[name]
	if !exists {
		// 列出可用的資料庫名稱，幫助診斷
		availableDatabases := make([]string, 0, len(dm.configs))
		for k := range dm.configs {
			availableDatabases = append(availableDatabases, k)
		}
		return nil, fmt.Errorf("資料庫 %s 未找到。可用的資料庫: %v", name, availableDatabases)
	}

	// 額外檢查資料庫連接是否有效
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("無法獲取資料庫連接: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("資料庫連接驗證失敗: %v", err)
	}

	return db, nil
}

// Close 關閉所有資料庫連接
func (dm *DatabaseManager) Close() error {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	var errs []error
	for name, db := range dm.configs {
		sqlDB, err := db.DB()
		if err != nil {
			errs = append(errs, fmt.Errorf("%s 資料庫關閉失敗: %v", name, err))
			continue
		}
		if err := sqlDB.Close(); err != nil {
			errs = append(errs, fmt.Errorf("%s 資料庫關閉失敗: %v", name, err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("關閉資料庫時出現錯誤: %v", errs)
	}

	return nil
}
