package logic

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/base/jwt"
	"tiktok/base/mymysql/tiktokdb"
	"tiktok/base/snowflake"
	"tiktok/service/commentsrv/models"
	"tiktok/service/usersrv/logic"
	"time"
)

const addComment = 1

var id int64

// CommentHandler 评论逻辑部分
func CommentHandler(c *gin.Context, p *io.ParamComment, data *io.CommentActionResponse) (err error) {
	// 构造评论返回值实例
	// 从参数token获取用户id,通过id获取用户信息
	data = new(io.CommentActionResponse)
	claim, err := jwt.ParseToken(p.Token)
	if err != nil {
		io.ResponseError(c, common.CodeNeedLogin)
		return
	}
	id = claim.UserID
	userInfoRep := new(io.UserInfoReq)
	userInfoRep.UserID = id
	userResp, err := logic.GetUserInfo(c, userInfoRep, claim)
	data.Comment.User = *userResp
	// 判断是删除评论还是添加评论
	if p.ActionType == addComment {
		err = AddComment(c, p, data)
		if err != nil {
			zap.L().Error("AddComment(c,p) failed ", zap.Error(err))
			return err
		}
	} else {
		data, err = DelComment(c, p)
		if err != nil {
			zap.L().Error("DelComment(c,p) failed ", zap.Error(err))
			return err
		}
	}
	return
}

func AddComment(c *gin.Context, p *io.ParamComment, data *io.CommentActionResponse) (err error) {
	// 生成评论id
	commentID := snowflake.GenID()
	// 构造实例对象
	t := time.Now()
	comment := &tiktokdb.Comment{
		CommentID:  commentID,
		VideoID:    p.VideoId,
		UserID:     id,
		Content:    p.CommentText,
		CreateTime: t,
	}
	data.Comment.Id = commentID
	data.Comment.Content = p.CommentText
	data.Comment.CreateDate = t.Format("01-02")
	// 保存到数据库
	return models.InsertComment(c, comment)
}

func DelComment(c *gin.Context, p *io.ParamComment) (data *io.CommentActionResponse, err error) {
	return
}
