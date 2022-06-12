package controller

import (
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/base/jwt"
	"tiktok/service/usersrv/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RegisterHandler 处理注册请求的函数
func RegisterHandler(c *gin.Context) {

	// 1. 获取参数和参数校验
	p := new(io.ParamRegister)
	p.Username = c.PostForm("username")
	p.Password = c.PostForm("password")
	//logger.Fdebug(fmt.Sprint(p))

	// 2. 服务调用
	// 目前是直接调用模块的 logic 功能
	user, err := logic.RegisterHandler(c, p)
	if err != nil {
		zap.L().Error("register failed", zap.Error(err))
		io.ResponseError(c, common.CodeRegisterFailed)
		return
	}
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		io.ResponseError(c, common.CodeTokenCreateErr)
		return
	}
	c.Set("userId", user.UserID)

	// 4. 返回成功响应
	io.ResponseSuccess4Login(c, token)
}

// LoginHandler 用户登录
func LoginHandler(c *gin.Context) {

	// 1. 获取参数和参数校验
	p := new(io.ParamLogin)
	p.Username = c.PostForm("username")
	p.Password = c.PostForm("password")
	//logger.Fdebug(fmt.Sprint(p))

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

//func UserInfo(c *gin.Context) {
//	token := c.Query("token")
//
//	if user, exist := usersLoginInfo[token]; exist {
//		c.JSON(http.StatusOK, UserResponse{
//			Response: io2.Response{StatusCode: 0},
//			User:     user,
//		})
//	} else {
//		c.JSON(http.StatusOK, UserResponse{
//			Response: io2.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
//		})
//	}
//}
