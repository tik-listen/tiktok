package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"sync/atomic"
	"tiktok/base/code"
	io2 "tiktok/base/io"
	"tiktok/base/mymysql/tiktokdb"
	"tiktok/service/usersrv/logic"
)

// RegisterHandler 处理注册请求的函数
func RegisterHandler(c *gin.Context) {

	// 1. 获取参数和参数校验
	p := new(io2.ParamRegister)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("register with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errors := err.(validator.ValidationErrors)
		if errors != nil {
			// 返回参数错误响应
			io2.ResponseError(c, code.CodeInvalidParam)
			return
		}
		return
	}

	// 2. 服务调用
	// 目前是直接调用模块的 logic 功能
	logic.RegisterHandler(c, p)
	// 3. 返回响应
	io2.ResponseSuccess(c, nil)
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	io2.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	io2.Response
	User tiktokdb.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: io2.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := tiktokdb.User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: io2.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: io2.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: io2.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: io2.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: io2.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
