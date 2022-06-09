package common

// ResCode int32
type ResCode int32

const (
	CodeSuccess          ResCode = 0
	CodeInvalidParam     ResCode = 1001
	CodeUserExist        ResCode = 1002
	CodeInvalidLoginInfo ResCode = 1004
	CodeServerBusy       ResCode = 1005
	CodeNeedLogin        ResCode = 1006
	CodeInvalidToken     ResCode = 1007
	CodeRegisterFailed   ResCode = 1008
	CodeTokenCreateErr   ResCode = 1009
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:          "success",
	CodeInvalidParam:     "请求参数错误",
	CodeUserExist:        "用户名已存在",
	CodeInvalidLoginInfo: "用户名或密码错误",
	CodeServerBusy:       "服务繁忙",
	CodeRegisterFailed:   "注册失败",
	CodeNeedLogin:        "需要登录",
	CodeInvalidToken:     "无效的token",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
