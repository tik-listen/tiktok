package logic

import (
	"github.com/gin-gonic/gin"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/base/jwt"
)

// CommentHandler 评论逻辑部分
func CommentHandler(c *gin.Context, p *io.ParamComment) (data *io.CommentActionResponse, err error) {
	// 构造评论返回值实例
	// 从参数token获取用户id,通过id获取用户信息
	data = new(io.CommentActionResponse)
	claim, err := jwt.ParseToken(p.Token)
	if err != nil {
		io.ResponseError(c, common.CodeNeedLogin)
		return
	}
	data.UserID = claim.UserID

	// 从参数获取评论内容

	// 数据库插入评论，返回时间

	// 构造完成，返回即可
	return
}
