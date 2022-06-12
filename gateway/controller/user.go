package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/base/jwt"
	"tiktok/service/usersrv/logic"
)

// RegisterHandler 处理注册请求的函数
func RegisterHandler(c *gin.Context) {

	// 1. 获取参数和参数校验
	// 绑定 Query 参数
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

	// 自动登录，获取 token
	user := &io.ParamLogin{
		Username: p.Username,
		Password: p.Password,
	}
	userId, token, err := logic.Login(user)
	if err != nil {
		io.ResponseError(c, common.CodeInvalidLoginInfo)
		return
	}
	c.Set("userId", userId)

	// 4. 返回成功响应
	io.ResponseSuccess4Login(c, token)
}

// LoginHandler 用户登录
func LoginHandler(c *gin.Context) {

	// 1. 获取参数和参数校验
	// 绑定 Query 参数
	p := new(io.ParamLogin)
	if err := c.ShouldBindWith(p, binding.Form); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		io.ResponseError(c, common.CodeInvalidParam)
		return
	}

	// 2. 服务调用
	// 目前是直接调用模块的 logic 功能
	userId, token, err := logic.Login(p)
	if err != nil {
		io.ResponseError(c, common.CodeInvalidLoginInfo)
		return
	}
	c.Set("userId", userId)
	// 3. 返回成功响应
	io.ResponseSuccess4Login(c, token)

}

func UserInfo(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(io.UserInfoReq)
	// 绑定 Query 参数
	if err := c.ShouldBindWith(p, binding.Form); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("get user info invalid param", zap.Error(err))
		io.ResponseError(c, common.CodeInvalidParam)
		return
	}
	// 登录校验
	claim, err := jwt.ParseToken(p.Token)
	if err != nil {
		io.ResponseError(c, common.CodeNeedLogin)
		return
	}
	//2. 服务调用
	//目前是直接调用模块的 logic 功能
	resp, err := logic.GetUserInfo(c, p, claim)
	if err != nil {
		io.ResponseError(c, common.CodeInvalidLoginInfo)
		return
	}

	// 3. 返回成功响应
	io.ResponseSuccessUserInfo(c, resp)

}
