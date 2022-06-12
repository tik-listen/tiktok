package controller

import (
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/service/favoritesrv/logic"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// VideoLikeAction 视频点赞操作
func VideoLikeAction(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(io.LikeActionReq)
	if err := c.ShouldBindWith(p, binding.Form); err != nil {
		zap.L().Error("VideoLikeAction with invalid param", zap.Error(err))
		errors := err.(validator.ValidationErrors)
		if errors != nil {
			io.ResponseError(c, common.CodeInvalidParam)
			return
		}
		return
	}

	// 2. 服务调用
	// 目前是直接调用模块的 logic 功能
	resp, err := logic.DealLikeAction(c, p)
	if err != nil {
		zap.L().Error("DealLikeAction failed", zap.Error(err))
		io.ResponseError(c, common.CodeRegisterFailed)
		return
	}
	io.RetResponse(c, resp)
}
func VideoLikeList(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(io.UserInfoReq)
	// 这里是针对 GET method 的操作
	if err := c.ShouldBindWith(p, binding.Form); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("VideoLikeList with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errors := err.(validator.ValidationErrors)
		if errors != nil {
			// 返回参数错误响应
			io.ResponseError(c, common.CodeInvalidParam)
			return
		}
		return
	}
	// 2. 服务调用
	// 目前是直接调用模块的 logic 功能
	resp, err := logic.GetFavoriteList(c, p)
	if err != nil {
		zap.L().Error("get Favorite list failed", zap.Error(err))
		io.ResponseError(c, common.CodeInvalidParam)
		return
	}
	io.RetFavoriteListReponse(c, resp)
}
