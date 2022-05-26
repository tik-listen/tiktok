package io

import "tiktok/base/mymysql/tiktokdb"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}


type FeedResponse struct {
	Response
	VideoList []tiktokdb.Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
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
