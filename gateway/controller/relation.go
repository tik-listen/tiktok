package controller

import (
	"fmt"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/service/usersrv/logic"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

// // RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	r := new(io.ParamRealation)
	if err := c.ShouldBindWith(r, binding.Form); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("RealationAction with invalid param", zap.Error(err))
		io.ResponseError(c, common.CodeInvalidParam)
		return
	}
	fmt.Println("debug begin")
	fmt.Println(r)
	fmt.Println("debug end")
	resp, err := logic.DealRelationAction(c, r)
	if err != nil {
		io.ResponseError(c, common.CodeInvalidParam)
		return
	}
	io.RetResponse(c, resp)
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	r := new(io.UserInfoReq)
	if err := c.ShouldBindWith(r, binding.Form); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("UserInfoReq with invalid param", zap.Error(err))
		io.ResponseError(c, common.CodeInvalidParam)
		return
	}
	fmt.Println("debug begin")
	fmt.Println(r)
	fmt.Println("debug end")
	resp, err := logic.FindFollweList(c, r)
	if err != nil {
		io.ResponseError(c, common.CodeInvalidParam)
		return
	}
	io.RetRelationResponse(c, resp)
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	r := new(io.UserInfoReq)
	if err := c.ShouldBindWith(r, binding.Form); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("UserInfoReq with invalid param", zap.Error(err))
		io.ResponseError(c, common.CodeInvalidParam)
		return
	}
	fmt.Println("debug begin")
	fmt.Println(r)
	fmt.Println("debug end")
	resp, err := logic.FindFollwerList(c, r)
	if err != nil {
		io.ResponseError(c, common.CodeInvalidParam)
		return
	}
	io.RetRelationResponse(c, resp)
}
