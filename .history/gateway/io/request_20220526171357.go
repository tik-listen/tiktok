package io

import (
	"errors"
	"tiktok/gateway/common"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("user not login")

// getCurrentUserID: Get the user id, who in login state
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(common.CtxUserIDKey)
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
