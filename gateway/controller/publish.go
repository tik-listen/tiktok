package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"path/filepath"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/base/jwt"
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
	finalFileName := fmt.Sprintf("%d_%s", MyClaims.UserID, data.Filename)
	saveFile := filepath.Join("./videosrv/", finalFileName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		zap.L().Error("video fail", zap.Error(err))
		io.ResponseError(c, common.CodeSaveFileErr)
		return
	}

}
