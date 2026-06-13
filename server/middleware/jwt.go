package middleware

import (
	"Helios/common"
	"Helios/dto"
	"Helios/model"
	"Helios/pkg/response"
	"Helios/pkg/utils"
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v3"
	"github.com/appleboy/gin-jwt/v3/core"
	"github.com/gin-gonic/gin"
	gojwt "github.com/golang-jwt/jwt/v5"
)

// JWT 认证中间件
func JWTAuth() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		// Realm 向用户显示的名称，必需，建议使用项目名称或公司名称
		Realm: common.Config.Auth.JWT.Realm,
		// 签名算法 - 可选值：HS256、HS384、HS512、RS256、RS384 或 RS512，默认为 HS256
		SigningAlgorithm: "HS256",
		// 用于签名的密钥，必需，建议使用随机生成的字符串，并且长度至少为 32 字节
		Key: []byte(common.Config.Auth.JWT.Key),
		// JWT Token的有效时长，可选，默认为一小时
		Timeout: time.Duration(common.Config.Auth.JWT.Timeout) * time.Second,
		// 用户登录校验回调函数
		Authenticator: authenticator,
		// 在登录中处理 Token 内容的回调函数，注意：这些数据未加密，注意数据安全
		PayloadFunc: payloadFunc,
		// 登录成功后响应的回调函数
		LoginResponse: loginResponse,
		// 登录失败后的响应回调函数
		Unauthorized: unauthorized,
		// 用户退出登录后的回调函数
		LogoutResponse: logoutResponse,
		// 已经登录后，解析每次请求中带的 Token 内容的回调函数
		IdentityHandler: identityHandler,
		// 已经登录后，验证每次请求中带的 Token 是否有效的回调函数
		Authorizer: authorizer,
		// 指定从哪里查找 Token
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// Token 前缀
		TokenHeadName: "Bearer",
	})
}

// 密码登录方式
func loginHandler(ctx *gin.Context) (id uint, err error) {
	// 获取请求体中的用户名和密码
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := utils.GetValidateErrorMessage(err, &req)
		common.SystemLog.Error("绑定请求体失败：", msg)
		return 0, errors.New(msg)
	}

	// 获取用户
	var user model.SystemUser
	if err = common.DB.Model(&model.SystemUser{}).
		Where("username = ?", req.Username).
		First(&user).Error; err != nil {
		common.SystemLog.Error("查询用户信息失败：", err.Error())
		return 0, err
	}
	if !utils.PasswordVerify(user.Password, req.Password) {
		common.SystemLog.Error("用户名或密码错误")
		return 0, errors.New("用户名或密码错误")
	}
	return user.Id, nil
}

// 钉钉扫码登录方式
func dingtalkLoginHandler(ctx *gin.Context) (id uint, err error) {
	// TODO: 处理钉钉扫码登录逻辑
	fmt.Println("钉钉扫码登录")
	return id, nil
}

// 飞书扫码登录方式
func feishuLoginHandler(ctx *gin.Context) (id uint, err error) {
	// TODO: 处理飞书扫码登录逻辑
	fmt.Println("飞书扫码登录")
	return id, nil
}

// 企业微信扫码登录方式
func wechatLoginHandler(ctx *gin.Context) (id uint, err error) {
	// TODO: 处理企业微信扫码登录逻辑
	fmt.Println("企业微信扫码登录")
	return id, nil
}

// 用户登录校验回调函数
func authenticator(ctx *gin.Context) (any, error) {
	// 获取请求的 uri 来区分不同的登录方式，同时为了防止有参数，需要去掉 ? 后面的部分，以及通过 / 分割路径来获取最后的部分
	uri := strings.Split(strings.Split(ctx.Request.RequestURI, "?")[0], "/")
	loginType := uri[len(uri)-1]
	var id uint
	var err error
	switch loginType {
	case "login":
		// 处理密码登录逻辑
		id, err = loginHandler(ctx)
	case "dingtalk":
		// 处理钉钉扫码登录逻辑
		id, err = dingtalkLoginHandler(ctx)
	case "feishu":
		// 处理飞书扫码登录逻辑
		id, err = feishuLoginHandler(ctx)
	case "wechat":
		// 处理企业微信扫码登录逻辑
		id, err = wechatLoginHandler(ctx)
	default:
		return nil, errors.New("登录认证方式暂不支持")
	}
	if err != nil {
		return nil, err
	}

	// 查询用户信息
	var user model.SystemUser
	if err = common.DB.Model(&model.SystemUser{}).
		Where("id = ?", id).
		Preload("SystemRole").
		First(&user).Error; err != nil {
		common.SystemLog.Error("查询用户信息失败：", err.Error())
		return nil, err
	}

	return &user, nil
}

// 登录中处理 Token 中内容的回调函数
func payloadFunc(data any) gojwt.MapClaims {
	if v, ok := data.(*model.SystemUser); ok {
		// 自定义需要将哪些数据放到 Token 中，注意不要放敏感数据，因为 Token 是可以被解密的
		return gojwt.MapClaims{
			"identityKey": v.Id,
			"id":          v.Id,
			"username":    v.Username,
			"nickname":    v.Nickname,
			"mobile":      v.Mobile,
			"email":       v.Email,
			"gender":      v.Gender,
			"avatar_url":  v.AvatarUrl,
			"status":      v.Status,
			"expire_at":   v.ExpireAt,
			"system_role": v.SystemRole.Name,
		}
	}
	return gojwt.MapClaims{}
}

// 登录成功后响应的回调函数
func loginResponse(ctx *gin.Context, token *core.Token) {
	resp := dto.LoginResponse{
		Token: token.AccessToken,
		// ExpireTime 是一个时间戳，表示 Token 的过期时间，可以根据需要格式化为 YYYY-MM-DD HH:MM:SS 格式
		ExpireTime: time.Unix(token.ExpiresAt, 0).Format(common.TIME_SECOND),
	}
	response.SuccessWithData(resp)
}

// 登录失败后的响应回调函数
func unauthorized(ctx *gin.Context, code int, message string) {
	response.FailureWithMessage(message)
}

// 用户退出登录后的回调函数
func logoutResponse(ctx *gin.Context) {
	// 解析请求中的 Token
	claims := jwt.ExtractClaims(ctx)
	id, _ := claims["id"].(float64)
	if uint(id) == 0 {
		response.FailureWithMessage("退出登录失败，未找到用户信息")
		return
	}
	// 清理用户相关登录数据
	response.SuccessWithMessage("退出登录成功")
}

// 已经登录后，解析每次请求中带的 Token 内容的回调函数
func identityHandler(ctx *gin.Context) any {
	claims := jwt.ExtractClaims(ctx)
	return &model.SystemUser{
		BaseModel: model.BaseModel{
			Id: uint(claims["id"].(float64)),
		},
	}
}

// 已经登录后，验证每次请求中带的 Token 是否有效的回调函数
func authorizer(ctx *gin.Context, data any) bool {
	// 新增和 Redis 中 Token 的对比逻辑
	if v, ok := data.(*model.SystemUser); ok && v.Id != 0 {
		return true
	}
	return false
}
