package dto

// 用户创建请求 - 必填字段用非指针类型
type SystemUserCreateRequest struct {
	Nickname     string `json:"nickname" binding:"required,min=2,max=50" msg:"昵称不能为空，长度必须在2-50之间"`
	Username     string `json:"username" binding:"required,min=2,max=50" msg:"用户名不能为空，长度必须在2-50之间"`
	Password     string `json:"password" binding:"required,min=6,max=128" msg:"密码不能为空，长度必须在6-128之间"`
	Mobile       string `json:"mobile" binding:"required,mobile" msg:"手机号格式不正确"`
	HideMobile   *uint  `json:"hide_mobile" binding:"required,oneof=0 1" msg:"隐藏手机号字段不能为空，且只能为 0（不隐藏）或 1（隐藏）"`
	Email        string `json:"email" binding:"required,email" msg:"邮箱格式不正确"`
	Gender       *uint  `json:"gender" binding:"required,oneof=0 1 2" msg:"性别字段不能为空，且只能为 0（未知）、1（男）或 2（女）"`
	AvatarUrl    string `json:"avatar_url" binding:"url" msg:"头像URL格式不正确"`
	Status       *uint  `json:"status" binding:"required,oneof=0 1" msg:"状态字段不能为空，且只能为 0（禁用）或 1（启用）"`
	ExpireAt     string `json:"expire_at" binding:"datetime=2006-01-02" msg:"过期时间格式不正确，必须为YYYY-MM-DD格式"`
	SystemRoleId uint   `json:"system_role_id" binding:"required" msg:"系统角色ID不能为空"`
}
