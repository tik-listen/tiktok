package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/source/apps/commentsrv/logic"
)

const MsgSuccess = "操作成功"
const MsgFailed = "操作失败"

// CommentHandler 评论动作
func CommentHandler(c *gin.Context) {
	// 获取参数
	p := new(io.ParamComment)
	if err := c.ShouldBindWith(p, binding.Form); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("comment with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errors := err.(validator.ValidationErrors)
		if errors != nil {
			// 返回参数错误响应
			io.ResponseError(c, common.CodeInvalidParam)
			return
		}
		return
	}
	// 逻辑处理
	data := new(io.CommentActionResponse)
	// 默认评论失败
	data.StatusCode = 1
	data.StatusMsg = MsgFailed
	err := logic.CommentHandler(c, p, data)

	if err != nil {
		zap.L().Error("logic.CommentHandler failed", zap.Error(err))
		io.ResponseError(c, common.CodeServerBusy)
		return
	}
	data.StatusCode = 0
	data.StatusMsg = MsgSuccess
	c.JSON(http.StatusOK, &data)
}

func GetCommentList(c *gin.Context) {
	// 获取参数
	p := new(io.ParmaCommentList)
	if err := c.ShouldBindWith(p, binding.Form); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("CommentList with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errors := err.(validator.ValidationErrors)
		if errors != nil {
			// 返回参数错误响应
			io.ResponseError(c, common.CodeInvalidParam)
			return
		}
		return
	}
	// 逻辑处理
	list, err := logic.GetCommentList(c, p)
	if err != nil {
		zap.L().Error("logic.GetCommentList(c, p) failed", zap.Error(err))
		io.ResponseError(c, common.CodeServerBusy)
		return
	}
	// 返回响应

	c.JSON(http.StatusOK, list)

}
