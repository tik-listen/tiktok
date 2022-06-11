package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/service/usersrv/logic"
)

//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"tiktok/base/io"
//)
//
//// FavoriteAction no practical effect, just check if token is valid
//func FavoriteAction(c *gin.Context) {
//	token := c.Query("token")
//
//	if _, exist := usersLoginInfo[token]; exist {
//		c.JSON(http.StatusOK, io.Response{StatusCode: 0})
//	} else {
//		c.JSON(http.StatusOK, io.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
//	}
//}
//
//// FavoriteList all users have same favorite video list
//func FavoriteList(c *gin.Context) {
//	c.JSON(http.StatusOK, VideoListResponse{
//		Response: io.Response{
//			StatusCode: 0,
//		},
//		VideoList: DemoVideos,
//	})
//}

// VideoLikeAction 视频点赞操作
func VideoLikeAction(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(io.ParamRegister)
	// 这里是针对 GET method 的操作
	if err := c.ShouldBindWith(p, binding.Form); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("register with invalid param", zap.Error(err))
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
	if err := logic.RegisterHandler(c, p); err != nil {
		zap.L().Error("register failed", zap.Error(err))
		io.ResponseError(c, common.CodeRegisterFailed)
		return
	}
}
