package common

// ResCode int32
type ResCode int32
type Action int32

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
	CodeVideoErr         ResCode = 1010
	CodeSaveFileErr      ResCode = 1011
	CodeVideoImFail      ResCode = 1012
	CodeGetVideoListErr  ResCode = 1013
)
const (
	Add    Action = 1
	Cancle Action = 2
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:          "success",
	CodeInvalidParam:     "请求参数错误",
	CodeUserExist:        "用户名已存在",
	CodeInvalidLoginInfo: "查不到该用户信息",
	CodeServerBusy:       "服务繁忙",
	CodeRegisterFailed:   "注册失败",
	CodeNeedLogin:        "需要登录",
	CodeInvalidToken:     "无效的token",
	CodeVideoErr:         "获取视频流失败",
	CodeSaveFileErr:      "保存文件失败",
	CodeVideoImFail:      "写库失败",
	CodeGetVideoListErr:  "获取视频列表失败",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
