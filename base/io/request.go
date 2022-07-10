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
	Username string `form:"username" binding:"required,min=4,max=32"`
	Password string `form:"password" binding:"required,min=6,max=32"`
}

// ParamLogin 登录请求参数，注意不要修改
type ParamLogin struct {
	Username string `form:"username" binding:"required,min=4,max=32"`
	Password string `form:"password" binding:"required,min=6,max=32"`
}

// UserInfoReq 用户信息请求参数
type UserInfoReq struct {
	UserID int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

// LikeActionReq 点赞请求
type LikeActionReq struct {
	Token      string        `form:"token" binding:"required"`
	VideoID    int64         `form:"video_id" binding:"required"`
	ActionType common.Action `form:"action_type" binding:"required,oneof=1 2"`
}

// ParamComment 评论相关参数
type ParamComment struct {
	Token       string `form:"token" binding:"required"`
	VideoId     int64  `form:"video_id" binding:"required"`
	ActionType  int64  `form:"action_type" binding:"required"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

// ParmaCommentList 拉取评论列表
type ParmaCommentList struct {
	Token   string `form:"token" binding:"required"`
	VideoId int64  `form:"video_id" binding:"required"`
}
type ParamRealation struct {
	Token      string        `form:"token" binding:"required"`
	ToUserID   int64         `form:"to_user_id" binding:"required"`
	ActionType common.Action `form:"action_type" binding:"required"`
}

// FeedRequest 请求视频流
type FeedRequest struct {
	LastTime int64  `json:"last_time"`
	Token    string `json:"token"`
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
