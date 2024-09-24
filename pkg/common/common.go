package common

const (
	UNAUTHORIZED_ACCESS = "无权限访问该服务"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
