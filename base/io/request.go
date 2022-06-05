package io

import (
	"errors"
	"tiktok/base/common"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("user not login")

// 请求参数写在这里面

// ParamRegister 注册请求参数，注意不要修改
type ParamRegister struct {
	Username string `json:"username" binding:"required,min=4,max=32"`
	Password string `json:"password" binding:"required,min=6,max=32"`
}

// ParamLogin 注册请求参数，注意不要修改
type ParamLogin struct {
	Username string `json:"username" binding:"required,min=4,max=32"`
	Password string `json:"password" binding:"required,min=6,max=32"`
}

// getCurrentUserID: Get the user id, who in login state
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(common.KCtxUserIDKey)
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
