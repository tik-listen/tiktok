package tiktokdb

import (
	"errors"
	"tiktok/base/jwt"
	"tiktok/base/mymysql"
	"tiktok/service/favoritesrv/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Video vedio 在redis的格式 video+时间（到秒）的Set存储视频id
type Video struct {
	VideoId       int64  `json:"id" db:"video_id"`
	UserId        int64  `json:"author" db:"user_id"`
	CoverUrl      string `json:"cover_url" db:"cover_url"`
	FavoriteCount int64  `json:"favorite_count" db:"favorite_count"`
	CommentCount  int64  `json:"comment_count" db:"comment_count"`
	IsFavorite    bool   `json:"is_favorite" db:"is_favorite"`
	Date          int64  `json:"date" db:"date"`
	Name          string `json:"name" db:"name"`
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
	if result := db.Table("video").Where("user_id = ? and name = ?", userid, name).Count(&count); result.Error != nil {
		return true
	}
	if count > 0 {
		return true
	}

	return false
}

// GetVideoListWithTime 获取某一时间之前的视频列表
func GetVideoListWithTime(c *gin.Context, now time.Time, token string) ([]Video, error) {
	db := mymysql.GetDB(c)
	res := make([]Video, 0)
	MyClaims, err := jwt.ParseToken(token)
	if err != nil {
		zap.L().Error("token is invalid", zap.Error(err))
		return nil, err
	}
	userID := MyClaims.UserID
	err = db.Table("video").Where("date < ?", now.Unix()).Limit(10).Find(&res).Error
	if err != nil {
		return nil, errors.New("MySQL ERR")
	}
	for idx, video := range res {
		favoites, err := models.FindFavoriteByVideoID(c, video.VideoId)
		if err != nil {
			return nil, err
		}
		res[idx].FavoriteCount = int64(len(favoites))
		for _, favoite := range favoites {
			if favoite.UserID == userID {
				res[idx].IsFavorite = true
			}
		}
	}
	for idx, video := range res {
		comments, err := FindCommentList(c, video.VideoId)
		if err != nil {
			return nil, err
		}
		res[idx].CommentCount = int64(len(comments))
	}

	return res, nil
}

// GetVideoListWithId  根据用户id获取投稿视频
func GetVideoListWithId(c *gin.Context, id int64) ([]Video, error) {
	db := mymysql.GetDB(c)
	var res []Video
	err := db.Table("video").Where("user_id = ?", id).Limit(10).Find(&res)
	if err.Error != nil {
		return nil, errors.New("MySQL ERR")
	}
	return res, nil
}

// GetVideoListWithVideoId  根据视频id获取信息
func GetVideoListWithVideoId(c *gin.Context, id int64) (Video, error) {
	db := mymysql.GetDB(c)
	var res Video
	err := db.Table("video").Where("video_id = ?", id).Find(&res)
	if err.Error != nil {
		return Video{}, errors.New("MySQL ERR")
	}
	return res, nil
}
