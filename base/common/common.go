package common

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	KCtxUserIDKey   = "userID"   // userId 上下文的 userId
	KCtxUserNameKey = "username" // username 上下文的 username
	Kmd5Secret      = "暂时先写在这"   // 用于用户信息加密
)

var (
	ErrorInvalidID         = errors.New("无效的ID")
	ErrorMysqlDbErr        = errors.New("MySQL 数据库错误")
	ErrorUserNotLogin      = errors.New("用户未登录")
	ErrorUserNotExist      = errors.New("数据库查询用户不存在")
	ErrorUserExist         = errors.New("用户已存在")
	ErrorDBError           = errors.New("数据库错误")
	ErrorInvalidPassword   = errors.New("密码错误，无效的查询")
	ErrorInvalid           = errors.New("参数错误")
	ErrorCommentNotExist   = errors.New("评论不存在")
	ErrorCommentNotEquUser = errors.New("删除评论用户与评论用户不同")
)

func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(KCtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
