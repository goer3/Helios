package dto

// 用户登录请求参数
type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// 用户登录响应参数
type LoginResponse struct {
	Token      string `json:"token"`
	ExpireTime string `json:"expire_time"`
}
