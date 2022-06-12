package logic

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/base/jwt"
	"tiktok/service/usersrv/models"
)

const addComment = 1

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
	data.Comment.User.UserID = claim.UserID
	_, err = models.FindOneUser(c, &data.Comment.User)

	// 判断是删除评论还是添加评论
	if p.ActionType == addComment {
		data, err = AddComment(c, p)
		if err != nil {
			zap.L().Error("AddComment(c,p) failed ", zap.Error(err))
			return nil, err
		}
	} else {
		data, err = DelComment(c, p)
		if err != nil {
			zap.L().Error("DelComment(c,p) failed ", zap.Error(err))
			return nil, err
		}
	}
	return
}

func AddComment(c *gin.Context, p *io.ParamComment) (data *io.CommentActionResponse, err error) {

	return
}

func DelComment(c *gin.Context, p *io.ParamComment) (data *io.CommentActionResponse, err error) {
	return
}
