package models

import (
	"github.com/gin-gonic/gin"
	"tiktok/base/mymysql/tiktokdb"
)

func InsertComment(c *gin.Context, comment *tiktokdb.Comment) error {
	return tiktokdb.InsertOneComment(c, comment)
}
