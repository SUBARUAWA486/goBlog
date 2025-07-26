package response

const (
	CodeSuccess       = 0    // 成功
	CodeInvalidParam  = 1001 // 参数错误
	CodeNotFound      = 1002 // 数据不存在
	CodeAuthFailed    = 1003 // 认证失败
	CodeForbidden     = 1004 // 无权限访问
	CodeServerError   = 2001 // 服务内部错误
	CodeTooManyReq    = 2002 // 请求频率过高
	CodeDBError       = 2003 // 数据库错误
)