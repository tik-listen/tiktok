package tiktokdb

import (
	"github.com/gin-gonic/gin"
	"tiktok/base/mymysql"
	"time"
)

type Comment struct {
	CommentID  int64     `db:"comment_id"`
	VideoID    int64     `db:"video_id"`
	UserID     int64     `db:"user_id"`
	Content    string    `db:"content"`
	CreateTime time.Time `db:"create_time"`
}

func InsertOneComment(c *gin.Context, comment *Comment) (err error) {
	db := mymysql.GetDB(c)
	db.Table("comments").Create(comment)
	return
}

func DeleteOneComment(c *gin.Context, cid int64) (err error) {
	db := mymysql.GetDB(c)
	db.Delete(&Comment{}, cid)
	return
}
func CheckCommentExist(c *gin.Context, cid int64) (flag bool, err error) {
	// 获取数据库连接
	db := mymysql.GetDB(c)

	var count int64
	if result := db.Table("comments").Where("comment_id = ?", cid).Count(&count); result.Error != nil {
		return true, result.Error
	}
	if count > 0 {
		return true, nil
	}

	return false, nil
}
func FindCommentList(c *gin.Context, vid int64) (clist []Comment, err error) {
	db := mymysql.GetDB(c)
	err = db.Table("comments").Where("video_id=?", vid).Find(&clist).Error
	return clist, err
}
