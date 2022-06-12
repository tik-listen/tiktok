package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/service/commentsrv/logic"
)

// import (
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// 	"tiktok/gateway/io"
// )

// CommentAction no practical effect, just check if token is valid
// func CommentAction(c *gin.Context) {
// 	token := c.Query("token")
// 	actionType := c.Query("action_type")

// 	if user, exist := usersLoginInfo[token]; exist {
// 		if actionType == "1" {
// 			text := c.Query("comment_text")
// 			c.JSON(http.StatusOK, io.CommentActionResponse{Response: Response{StatusCode: 0},
// 				Comment: Comment{
// 					Id:         1,
// 					User:       user,
// 					Content:    text,
// 					CreateDate: "05-01",
// 				}})
// 			return
// 		}
// 		c.JSON(http.StatusOK, Response{StatusCode: 0})
// 	} else {
// 		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
// 	}
// }

// // CommentList all videos have same demo comment list
// func CommentList(c *gin.Context) {
// 	c.JSON(http.StatusOK, io.CommentListResponse{
// 		Response:    Response{StatusCode: 0},
// 		CommentList: DemoComments,
// 	})
// }

const MsgSuccess = "操作成功"
const MsgFailed = "操作失败"

// CommentHandler 新增评论
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
