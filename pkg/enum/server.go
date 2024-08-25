package enum

// ----------------全局状态码------------------
const (
	HttpSuccess     = 200   // 成功
	HttpFail        = 400   // 失败
	HttpError       = 500   // 错误
	InvalidArgument = 10000 // 非法参数
	TokenFailure    = 10001 // 登录令牌失效
)

// ----------------特殊常量--------------------
const (
	TimeDay               string = "2006-01-02"
	TimeMinutesAndSeconds string = "2006-01-02 15:04:05"
	JwtKey                string = "blog:jwt"
)
