package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"path/filepath"
	"strconv"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/base/jwt"
	"tiktok/base/mymysql/tiktokdb"
	"tiktok/base/snowflake"
	"tiktok/source/apps/publishsrv"
)

func PublishActionHandler(c *gin.Context) {
	//解析token
	token := c.PostForm("token")
	MyClaims, err := jwt.ParseToken(token)
	if err != nil {
		zap.L().Error("token is invalid", zap.Error(err))
		io.ResponseError(c, common.CodeTokenCreateErr)
		return
	}
	//拿到文件流
	data, err := c.FormFile("data")
	if err != nil {
		zap.L().Error("video fail", zap.Error(err))
		io.ResponseError(c, common.CodeVideoErr)
		return
	}
	name := data.Filename
	videoId := snowflake.GenID()
	//更新缓存和数据库
	err = publishsrv.SaveVideoIm(name, MyClaims.UserID, videoId, c)
	if err != nil {
		zap.L().Error("sql err", zap.Error(err))
		io.ResponseError(c, common.CodeVideoImFail)
		return
	}
	//写入文件
	saveFile := filepath.Join("./videosrv/", strconv.FormatInt(videoId, 10))
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		zap.L().Error("video fail", zap.Error(err))
		io.ResponseError(c, common.CodeSaveFileErr)
		return
	}
	io.ResponseSuccessVideoAction(c)
}
func PublishListHandler(c *gin.Context) {
	token := c.PostForm("token")
	_, err := jwt.ParseToken(token)
	if err != nil {
		zap.L().Error("token is invalid", zap.Error(err))
		io.ResponseError(c, common.CodeTokenCreateErr)
		return
	}
	userId := c.PostForm("id")
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		zap.L().Error("get id err", zap.Error(err))
		io.ResponseError(c, common.CodeInvalidParam)
		return
	}
	videoList, err := tiktokdb.GetVideoListWithId(c, id)
	if err != nil {
		zap.L().Error("sql err", zap.Error(err))
		io.ResponseError(c, common.CodeVideoImFail)
		return
	}
	io.ResponseSuccessPublishList(c, videoList)
}
