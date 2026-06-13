package response

// 请求结果状态码
const (
	OK                  = 200
	BadRequest          = 400
	Forbidden           = 403
	NotFound            = 404
	MethodNotAllowed    = 405
	InternalServerError = 500
	BadGateway          = 502
	ServiceUnavailable  = 503
	GatewayTimeout      = 504
)

// 请求结果状态码对应的信息描述
const (
	OKMessage                  = "请求成功"
	BadRequestMessage          = "请求失败"
	ForbiddenMessage           = "无权限访问该资源"
	NotFoundMessage            = "未找到相关资源"
	MethodNotAllowedMessage    = "请求方法不允许"
	InternalServerErrorMessage = "服务器错误"
	BadGatewayMessage          = "网关错误"
	ServiceUnavailableMessage  = "服务不可达"
	GatewayTimeoutMessage      = "网关超时"
)

// 请求结果状态码和描述信息做绑定
var ResponseMessage = map[int]string{
	OK:                  OKMessage,
	BadRequest:          BadRequestMessage,
	Forbidden:           ForbiddenMessage,
	NotFound:            NotFoundMessage,
	MethodNotAllowed:    MethodNotAllowedMessage,
	InternalServerError: InternalServerErrorMessage,
	BadGateway:          BadGatewayMessage,
	ServiceUnavailable:  ServiceUnavailableMessage,
	GatewayTimeout:      GatewayTimeoutMessage,
}

// 统一响应格式
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// 没有数据的时候响应空的数据
var empty = map[string]any{}

// 响应方法
func Result(code int, message string, data any) {
	panic(Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// 成功响应
func Success() {
	Result(OK, ResponseMessage[OK], empty)
}

// 成功响应，带消息
func SuccessWithMessage(message string) {
	Result(OK, message, empty)
}

// 成功响应，带数据
func SuccessWithData(data any) {
	Result(OK, ResponseMessage[OK], data)
}

// 失败响应
func Failure() {
	Result(BadRequest, ResponseMessage[BadRequest], empty)
}

// 失败响应，带状态码
func FailureWithCode(code int) {
	Result(code, ResponseMessage[code], empty)
}

// 失败响应，带消息
func FailureWithMessage(message string) {
	Result(BadRequest, message, empty)
}

// 失败响应，带状态码和消息
func FailureWithCodeAndMessage(code int, message string) {
	Result(code, message, empty)
}
