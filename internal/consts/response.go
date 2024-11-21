package consts

const (
	// 成功类错误码
	Success  = 20000 // 操作成功
	Created  = 20100 // 资源创建成功
	Accepted = 20200 // 请求已接受但未处理

	// 客户端错误类
	BadRequest       = 40000 // 请求错误（缺少参数、非法参数等）
	Unauthorized     = 40100 // 未授权（需要登录）
	Forbidden        = 40300 // 禁止访问（无权限）
	NotFound         = 40400 // 资源未找到
	MethodNotAllowed = 40500 // 请求方法不允许
	Timeout          = 40800 // 请求超时
	LoginFailed      = 41000 // 登录失败

	// 服务器错误类
	InternalServerError = 50000 // 服务器内部错误
	NotImplemented      = 50100 // 未实现
	ServiceUnavailable  = 50300 // 服务不可用
	GatewayTimeout      = 50400 // 网关超时

	// 数据相关错误
	DataNotFound = 60000 // 数据未找到

	InvalidData   = 60100 // 数据无效
	DuplicateData = 60200 // 数据重复

	// 验证相关错误
	ValidationError = 70000 // 数据验证失败
	InvalidToken    = 70100 // 无效的 Token
	ExpiredToken    = 70200 // Token 已过期

	// 自定义错误类
	CustomError = 80000 // 自定义错误

	// 业务相关错误
	OperationFailed   = 90000 // 业务操作失败
	InsufficientFunds = 90100 // 余额不足
)

var responseMap = map[int]string{
	// 成功类错误码
	Success:  "success",
	Created:  "资源创建成功",
	Accepted: "请求已接受但未处理",

	// 客户端错误类
	BadRequest:       "请求错误（缺少参数、非法参数等）",
	Unauthorized:     "未授权（需要登录）",
	Forbidden:        "禁止访问（无权限）",
	NotFound:         "资源未找到",
	MethodNotAllowed: "请求方法不允许",
	Timeout:          "请求超时",
	LoginFailed:      "登录失败",

	// 服务器错误类
	InternalServerError: "服务器内部错误",
	NotImplemented:      "未实现",
	ServiceUnavailable:  "服务不可用",
	GatewayTimeout:      "网关超时",

	// 数据相关错误
	DataNotFound:  "数据未找到",
	InvalidData:   "数据无效",
	DuplicateData: "数据重复",

	// 验证相关错误
	ValidationError: "数据验证失败",
	InvalidToken:    "无效的 Token",
	ExpiredToken:    "Token 已过期",

	// 自定义错误类
	CustomError: "自定义错误",

	// 业务相关错误
	OperationFailed:   "业务操作失败",
	InsufficientFunds: "余额不足",
}

func GetResponseMsg(code int) string {
	return responseMap[code]
}
