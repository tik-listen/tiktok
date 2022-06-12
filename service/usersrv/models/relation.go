package models

import (
	"tiktok/base/common"
	"tiktok/base/mymysql"

	"github.com/gin-gonic/gin"
)

type Relation struct {
	RelationID int64 `db:"relation_id"`
	UserID     int64 `db:"user_id"`
	ToUserID   int64 `db:"to_user_id"`
}

func InsertRelation(c *gin.Context, relation Relation) error {
	db := mymysql.GetDB(c)
	err := db.Table("relation").Create(relation).Error
	if err != nil {
		return common.ErrorMysqlDbErr
	}
	return nil
}

func DeleteRelation(c *gin.Context, relation Relation) error {
	db := mymysql.GetDB(c)
	err := db.Table("relation").Delete("user_id=? and to_user_id =?", relation.UserID, relation.ToUserID).Error
	if err != nil {
		return common.ErrorDBError
	}
	return nil
}

func FindUserFans(c *gin.Context, userid int64) (relations []Relation, err error) {
	db := mymysql.GetDB(c)
	err = db.Table("relation").Where("to_user_id=?", userid).Find(relations).Error
	return relations, err
}

func FindUserStar(c *gin.Context, userid int64) (relations []Relation, err error) {
	db := mymysql.GetDB(c)
	err = db.Table("relation").Where("user_id=?", userid).Find(relations).Error
	return relations, err
}
