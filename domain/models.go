package domain

// UserWithRoles 擴展用戶模型，包含角色
// 原 User strcut by design，待修改
type UserWithRoles struct {
	UID      string `json:"uid"`
	Username string `json:"username"`
	Roles    []Role `json:"roles"`
}

// Role 角色模型
type Role struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Permissions []Permission `json:"permissions"`
}

// Permission 權限模型
type Permission struct {
	ID          string `json:"id"`
	Resource    string `json:"resource"`
	Action      string `json:"action"`
	Description string `json:"description"`
}
