package initialize

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 手机号验证
func ValidateMobile(fl validator.FieldLevel) bool {
	// 手机号正则校验
	r := regexp.MustCompile(`^1[3456789]\d{9}$`)
	v := fl.Field().String()
	return r.MatchString(v)
}

// 初始化 gin 验证器
func Validator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证函数
		_ = v.RegisterValidation("mobile", ValidateMobile)
	}
}
