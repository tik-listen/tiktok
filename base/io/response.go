package io

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/base/code"
	"tiktok/base/mymysql/tiktokdb"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type FeedResponse struct {
	Response
	VideoList []tiktokdb.Video `json:"video_list,omitempty"`
	NextTime  int64            `json:"next_time,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []tiktokdb.Video `json:"video_list"`
}

type CommentListResponse struct {
	Response
	CommentList []tiktokdb.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment tiktokdb.Comment `json:"comment,omitempty"`
}

type UserListResponse struct {
	Response
	UserList []tiktokdb.User `json:"user_list"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User tiktokdb.User `json:"user"`
}

// ResponseData 响应的 Data
type ResponseData struct {
	Code code.ResCode `json:"code"`
	Msg  interface{}  `json:"msg"`
	Data interface{}  `json:"data,omitempty"`
}

// ResponseError 响应错误
func ResponseError(c *gin.Context, code code.ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// ResponseErrorWithMsg 响应错误附上 msg 信息
func ResponseErrorWithMsg(c *gin.Context, code code.ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccess 响应成功
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code.CodeSuccess,
		Msg:  code.CodeSuccess.Msg(),
		Data: data,
	})
}
