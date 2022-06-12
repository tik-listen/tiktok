package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/base/jwt"
	"tiktok/base/mymysql/tiktokdb"
	"tiktok/base/myredis"
	"time"
)

func FeedHandler(c *gin.Context) {
	lastTime := c.PostForm("last_time")
	token := c.PostForm("token")
	if lastTime == "" {
		if token != "" {
			_, err := jwt.ParseToken(token)
			if err != nil {
				zap.L().Error("token is invalid", zap.Error(err))
				io.ResponseError(c, common.CodeTokenCreateErr)
				return
			}
			//添加个性推荐的参数并返回
			//io.ResponseSuccessVideoList(c,个性化推荐列表)
		}
		data, err := myredis.GetVideoList()
		if err != nil {
			data, err := tiktokdb.GetVideoListWithTime(c, time.Now())
			if err != nil {
				io.ResponseError(c, common.CodeGetVideoListErr)
			}
			io.ResponseSuccessVideoList(c, data)
		}

		var res []tiktokdb.Video
		for i := 0; i < 10; i++ {
			b := []byte(data[i])
			_ = json.Unmarshal(b, &res[i])
		}
		io.ResponseSuccessVideoList(c, res)
		return
	} else {
		now := c.Query("last_time")
		if token != "" {
			_, err := jwt.ParseToken(token)
			if err != nil {
				zap.L().Error("token is invalid", zap.Error(err))
				io.ResponseError(c, common.CodeTokenCreateErr)
				return
			}
			//添加个性推荐的参数并返回
			//io.ResponseSuccessVideoList(c,个性化推荐列表)
		}
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", now, time.Local)
		data, err := tiktokdb.GetVideoListWithTime(c, t)
		if err != nil {
			io.ResponseError(c, common.CodeGetVideoListErr)
		}
		io.ResponseSuccessVideoList(c, data)
		return
	}

}
