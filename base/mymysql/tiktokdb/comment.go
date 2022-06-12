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
