package io

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/base/common"
	"tiktok/base/mymysql/tiktokdb"
)

type Response struct {
	StatusCode common.ResCode `json:"status_code"`
	StatusMsg  string         `json:"status_msg,omitempty"`
}
type VideoRes struct {
	Id            int64         `json:"id,omitempty" db:"video_id"`
	User          tiktokdb.User `json:"author" db:"user_id"`
	PlayUrl       string        `json:"play_url"`
	CoverUrl      string        `json:"cover_url" db:"cover_url"`
	FavoriteCount int64         `json:"favorite_count" db:"favorite_count"`
	CommentCount  int64         `json:"comment_count" db:"comment_count"`
	IsFavorite    bool          `json:"is_favorite" db:"is_favorite"`
	Name          string        `json:"title" db:"name"`
}
type VideoPublishList struct {
	Response
	VideoList []VideoRes `json:"video_list"`
}
type FeedResponse struct {
	Response
	VideoList []VideoRes `json:"video_list,omitempty"`
	NextTime  int64      `json:"next_time,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []VideoRes `json:"video_list"`
	Time      int64      `json:"next_time"`
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

// ResponseSuccess4Login 登录成功
func ResponseSuccess4Login(c *gin.Context, token string) {
	userId, _ := c.Get("userId")
	c.JSON(http.StatusOK, &UserLoginResponse{
		Response: Response{common.CodeSuccess, common.CodeSuccess.Msg()},
		UserId:   userId.(int64),
		Token:    token,
	})
}

// ResponseSuccessVideoAction 投稿上传视频成功
func ResponseSuccessVideoAction(c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		StatusCode: 0,
		StatusMsg:  "上传成功",
	})
}

// ResponseSuccessVideoList 获取时间排序的视频列表成功
func ResponseSuccessVideoList(c *gin.Context, videoList []tiktokdb.Video) {
	n := len(videoList)
	var res = make([]VideoRes, n)
	time := int64(0)
	for i := 0; i < n; i++ {
		res[i].Id = videoList[i].VideoId
		res[i].User, _ = tiktokdb.GetOneUserWithId(c, videoList[i].UserId)
		res[i].PlayUrl = "http://82.157.141.199/" + strconv.FormatInt(videoList[i].VideoId, 10) + ".mp4"
		res[i].FavoriteCount = videoList[i].FavoriteCount
		res[i].CommentCount = videoList[i].FavoriteCount
		res[i].IsFavorite = videoList[i].IsFavorite
		res[i].Name = videoList[i].Name
		if videoList[i].Date > time {
			time = videoList[i].Date
		}
	}
	c.JSON(http.StatusOK, &VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "获取成功",
		},
		VideoList: res,
		Time:      time,
	})
}

// ResponseSuccessPublishList 获取投稿视频成功
func ResponseSuccessPublishList(c *gin.Context, videoList []tiktokdb.Video) {
	n := len(videoList)
	var res = make([]VideoRes, n)
	for i := 0; i < n; i++ {
		res[i].Id = videoList[i].VideoId
		res[i].User, _ = tiktokdb.GetOneUserWithId(c, videoList[i].UserId)
		res[i].PlayUrl = "http://82.157.141.199/" + strconv.FormatInt(videoList[i].VideoId, 10) + ".mp4"
		res[i].FavoriteCount = videoList[i].FavoriteCount
		res[i].CommentCount = videoList[i].FavoriteCount
		res[i].IsFavorite = videoList[i].IsFavorite
		res[i].Name = videoList[i].Name

	}
	c.JSON(http.StatusOK, &VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "获取成功",
		},
		VideoList: res,
	})
}
