package publishsrv

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"tiktok/base/mymysql/tiktokdb"
	"tiktok/base/myredis"
	"time"
)

func SaveVideoIm(name string, userId int64, videoId int64, c *gin.Context) error {
	if tiktokdb.CheckVideoExist(c, name, userId) {
		return errors.New("video has exist")
	}
	video := tiktokdb.Video{
		VideoId:       videoId,
		UserId:        userId,
		CoverUrl:      "",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Date:          time.Now().Unix(),
		Name:          name,
	}
	err := tiktokdb.InsertVideo(video, c)
	if err != nil {
		return err
	}
	b, _ := json.Marshal(video)
	myredis.SetListForVideoList(string(b))
	return nil
}
