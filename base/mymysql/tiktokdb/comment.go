package tiktokdb

import (
	"github.com/gin-gonic/gin"
	"tiktok/base/mymysql"
)

type Comment struct {
	CommentID  int64  `db:"comment_id"`
	VideoID    int64  `db:"video_id"`
	UserID     int64  `db:"user_id"`
	CreateTime int64  `db:"create_time"`
	Content    string `db:"content"`
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

// CheckCommentExist 根据cid判断评论是否存在
func CheckCommentExist(c *gin.Context, cid int64) (flag bool, err error) {
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

// GetUIDbyCID 通过cid查询uid
func GetUIDbyCID(c *gin.Context, cid int64) (uid int64, err error) {
	db := mymysql.GetDB(c)
	var comment Comment
	err = db.Table("comments").Where("comment_id = ?", cid).Find(&comment).Error
	uid = comment.UserID
	return
}

func FindCommentList(c *gin.Context, vid int64) (clist []Comment, err error) {
	db := mymysql.GetDB(c)
	err = db.Table("comments").Where("video_id=?", vid).Find(&clist).Error
	return clist, err
}

// GetCommentCount 查询视频的评论数
func GetCommentCount(c *gin.Context, vid int64) (count int64, err error) {
	db := mymysql.GetDB(c)
	err = db.Table("comments").Where("video_id=?", vid).Count(&count).Error
	return
}
