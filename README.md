# 遊戲平台 RBAC 權限管理系統

## 1. 項目概述
原，本項目旨在開發一個獨立的、基於HTTP調用的RBAC（基於角色的訪問控制）服務，為微服務架構提供統一的權限管理。該服務作為權限中心，負責用戶、角色、權限的管理與授權驗證。
本項目旨在開發一個獨立的、基於HTTP調用的 api server。

## 2. API端點

### 2.1 用戶管理
- [x] `POST /v1/users` - 創建用戶
- [x] `GET /v1/users/{id}` - 獲取指定用戶
- [ ] - `GET /v1/users` - 查詢用戶列表
- `PUT /v1/users/{id}` - 更新用戶信息
- `DELETE /v1/users/{id}` - 刪除用戶

### 2.2 角色管理
- `POST /v1/roles` - 創建角色
- `GET /v1/roles` - 查詢角色列表
- `GET /v1/roles/{id}` - 獲取指定角色
- `PUT /v1/roles/{id}` - 更新角色
- `DELETE /v1/roles/{id}` - 刪除角色

### 2.3 權限管理
- `POST /v1/permissions` - 創建權限
- `GET /v1/permissions` - 查詢權限列表
- `GET /v1/permissions/{id}` - 獲取指定權限
- `PUT /v1/permissions/{id}` - 更新權限
- `DELETE /v1/permissions/{id}` - 刪除權限

### 2.4 關聯管理
- `POST /v1/users/{id}/roles` - 為用戶分配角色
- `DELETE /v1/users/{id}/roles/{roleId}` - 移除用戶的角色
- `GET /v1/users/{id}/permissions` - 獲取用戶所有權限
- `POST /v1/roles/{id}/permissions` - 為角色分配權限
- `DELETE /v1/roles/{id}/permissions/{permId}` - 移除角色的權限

### 2.5 認證和授權
- [x] `POST /v1/auth/login` - 登入
- [x] `POST /v1/auth/login` - 登出 
- `POST /v1/auth/authorize` - 權限驗證 -> pending
- `POST /v1/auth/refresh` - 刷新令牌
- `POST /v1/auth/revoke` - 取消授權jwt
- `POST /v1/auth/batch-revoke` - 批量取消授權jwt

### 2.6 審計日誌
- `GET /v1/audit-logs` - 查詢審計日誌

## 3. 中介層
### 3.1 jwt 驗證
- [x] 檢查 token 是否為空
- [x] 檢查 token 是否過期
- [x] 檢查 token 是否與數據庫一致
### 3.2 錯誤攔截與統一處理
