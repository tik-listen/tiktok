package io

import (
	"net/http"
	"tiktok/base/common"
	"tiktok/base/mymysql/tiktokdb"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode common.ResCode `json:"status_code"`
	StatusMsg  string         `json:"status_msg,omitempty"`
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

// CommentActionResponse 评论返回值
type CommentActionResponse struct {
	UserInfoReq
	Comment tiktokdb.Comment `json:"comment,omitempty"`
}

type UserListResponse struct {
	Response
	UserList []tiktokdb.User `json:"user_list"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User tiktokdb.User `json:"user"`
}
type RelationResponse struct {
	Response
	UserList []UserInfoResp `json:"user_list"`
}

// ResponseData 通用的响应内容
type ResponseData struct {
	Response
	Msg  interface{} `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// ResponseError 响应错误
func ResponseError(c *gin.Context, code common.ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Response: Response{code, code.Msg()},
	})
}

// UserInfoResp 用户信息返回值
type UserInfoResp struct {
	Response
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// ResponseSuccess4Login 登录成功
func ResponseSuccess4Login(c *gin.Context, token string) {
	userId, _ := c.Get("userId")
	c.JSON(http.StatusOK, &UserLoginResponse{
		Response: Response{common.CodeSuccess, common.CodeSuccess.Msg()},
		UserId:   userId.(int64),
		Token:    token,
	})
}

// ResponseSuccessUserInfo 返回用户信息
func ResponseSuccessUserInfo(c *gin.Context, resp *UserInfoResp) {
	c.JSON(http.StatusOK, resp)
}
func RetResponse(c *gin.Context, resp *Response) {
	c.JSON(http.StatusOK, resp)
}
func RetRelationResponse(c *gin.Context, resp *RelationResponse) {
	c.JSON(http.StatusOK, resp)
}
