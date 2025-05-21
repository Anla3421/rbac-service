package domain

// Response 通用的API響應結構
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// LoginResponse 登錄響應
type LoginResponse struct {
	User  string `json:"user"`
	Token string `json:"token"`
}

// AuthorizeResponse 授權響應
type AuthorizeResponse struct {
	Authorized   bool  `json:"authorized"`
	NeedsRefresh bool  `json:"needsRefresh,omitempty"`
	ExpiresIn    int64 `json:"expiresIn,omitempty"`
}

// NewResponse 創建一個成功的響應
func NewResponse(message string, data interface{}) Response {
	return Response{
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse 創建一個錯誤響應
func NewErrorResponse(message string, err string) Response {
	return Response{
		Message: message,
		Error:   err,
	}
}
