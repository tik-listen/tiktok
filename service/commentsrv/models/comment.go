package models

import (
	"github.com/gin-gonic/gin"
	"tiktok/base/mymysql/tiktokdb"
)

func InsertComment(c *gin.Context, comment *tiktokdb.Comment) error {
	return tiktokdb.InsertOneComment(c, comment)
}

func DeleteComment(c *gin.Context, cid int64) error {
	return tiktokdb.DeleteOneComment(c, cid)
}

func CheckCommentExist(c *gin.Context, cid int64) (flag bool, err error) {
	return tiktokdb.CheckCommentExist(c, cid)
}
func FindCommentList(c *gin.Context, vid int64) (clist []tiktokdb.Comment, err error) {
	return tiktokdb.FindCommentList(c, vid)
}
