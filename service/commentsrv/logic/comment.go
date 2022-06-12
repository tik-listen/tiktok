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
	claim, err := jwt.ParseToken(p.Token)
	if err != nil {
		io.ResponseError(c, common.CodeNeedLogin)
		return
	}
	id = claim.UserID
	userInfoRep := new(io.UserInfoReq)
	userInfoRep.UserID = id
	userResp, err := logic.GetUserInfo(c, userInfoRep, claim)
	if err != nil {
		zap.L().Error("logic.GetUserInfo(c, userInfoRep, claim) failed", zap.Error(err))
		return
	}
	data.Comment.User = *userResp
	// 判断是删除评论还是添加评论
	if p.ActionType == addComment {
		err = AddComment(c, p, data)
		if err != nil {
			zap.L().Error("AddComment(c,p) failed ", zap.Error(err))
			return err
		}
	} else {
		err = DelComment(c, p.CommentId)
		if err != nil {
			zap.L().Error("DelComment(c,p) failed ", zap.Error(err))
			return err
		}
	}
	return
}

// AddComment 添加评论
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

// DelComment 删除评论
func DelComment(c *gin.Context, cid int64) (err error) {
	// 判断是否存在
	// 1.判断用户存不存在
	flag, err := models.CheckCommentExist(c, cid)
	if err != nil {
		return common.ErrorMysqlDbErr
	}
	if !flag {
		return common.ErrorCommentNotExist
	}

	// 判断cid是否等于用户id
	if cid != id {
		return common.ErrorCommentNotEquUser
	}
	return models.DeleteComment(c, cid)
}

func GetCommentList(c *gin.Context, p *io.ParmaCommentList) (list *io.CommentListResponse, err error) {
	clist, err := models.FindCommentList(c, p.VideoId)
	list = new(io.CommentListResponse)
	if err != nil {
		return list, err
	}
	claim, err := jwt.ParseToken(p.Token)
	if err != nil {
		io.ResponseError(c, common.CodeNeedLogin)
		return
	}
	list.Response.StatusCode = common.CodeSuccess
	list.Response.StatusMsg = "success"
	list.CommentList = make([]io.Comment, 0)
	for _, comment := range clist {
		tempComment := new(io.Comment)
		tempComment.Id = comment.CommentID
		tempComment.Content = comment.Content
		tempComment.CreateDate = comment.CreateTime.Format("01-02")

		userInfoRep := new(io.UserInfoReq)
		userInfoRep.UserID = comment.UserID
		userResp, err := logic.GetUserInfo(c, userInfoRep, claim)
		if err != nil {
			zap.L().Error("logic.GetUserInfo(c, userInfoRep, claim) failed", zap.Error(err))
			return nil, err
		}
		tempComment.User = *userResp
		list.CommentList = append(list.CommentList, *tempComment)
	}
	return
}
