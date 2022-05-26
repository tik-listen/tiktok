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
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}



type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}
