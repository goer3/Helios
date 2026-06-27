package dto

// 用户登录请求参数
type LoginRequest struct {
	Account  string `json:"account" form:"account" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// 用户登录响应参数
type LoginResponse struct {
	Token      string `json:"token"`
	ExpireTime string `json:"expire_time"`
}
