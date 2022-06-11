package tiktokdb

import (
	"errors"
	"github.com/gin-gonic/gin"
	"tiktok/base/mymysql"
	"time"
)

// Video vedio 在redis的格式 video+时间（到秒）的Set存储视频id
type Video struct {
	Id            int64     `json:"id,omitempty" db:"video_id"`
	UserId        int64     `json:"author" db:"user_id"`
	CoverUrl      string    `json:"cover_url" db:"cover_url"`
	FavoriteCount int64     `json:"favorite_count" db:"favorite_count"`
	CommentCount  int64     `json:"comment_count" db:"comment_count"`
	IsFavorite    bool      `json:"is_favorite" db:"is_favorite"`
	Date          time.Time `json:"date" db:"date"`
	Name          string    `json:"name" db:"name"`
}

// InsertVideo 插入video结构体
func InsertVideo(video Video, c *gin.Context) error {
	db := mymysql.GetDB(c)
	return db.Table("video").Create(video).Error
}

// CheckVideoExist 根据视频名，用户id，检查video是否存在
func CheckVideoExist(ctx *gin.Context, name string, userid int64) bool {
	db := mymysql.GetDB(ctx)
	var count int64
	if result := db.Table("video").Where("user_id = ?", userid, "name = ?", name).Count(&count); result.Error != nil {
		return true
	}
	if count > 0 {
		return true
	}

	return false
}

// GetVideoListWithTime 获取某一时间之前的视频列表
func GetVideoListWithTime(c *gin.Context, now time.Time) ([]Video, error) {
	db := mymysql.GetDB(c)
	var res []Video
	err := db.Table("video").Where("data < ", now).Limit(10).Find(&res)
	if err != nil {
		return nil, errors.New("MySQL ERR")
	}
	return res, nil
}
