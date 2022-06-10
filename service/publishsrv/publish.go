package publishsrv

import (
	"errors"
	"github.com/gin-gonic/gin"
	"tiktok/base/mymysql/tiktokdb"
	"tiktok/base/myredis"
)

func SaveVideoIm(name string, userId int64, videoId int64, c *gin.Context) error {
	if tiktokdb.CheckVideoExist(c, videoId) {
		return errors.New("video has exist")
	}

	err := tiktokdb.InsertVideo(name, userId, videoId, c)
	if err != nil {
		return err
	}
	return nil
}
