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

// CountUserFans 获取用户的粉丝数
func CountUserFans(c *gin.Context, userid int64) (int64, error) {
	db := mymysql.GetDB(c)
	var count int64 = 0
	//select count(*) from relation where to_user_id = user_id
	err := db.Table("relation").Where("to_user_id=?", userid).Count(&count).Error
	return count, err
}

// CountUserStar 获取用户的关注数
func CountUserStar(c *gin.Context, userid int64) (int64, error) {
	db := mymysql.GetDB(c)
	var count int64 = 0
	//select count(*) from relation where user_id = user_id
	err := db.Table("relation").Where("user_id=?", userid).Count(&count).Error
	return count, err
}
func IsFans(c *gin.Context, userid, touserid int64) (bool, error) {
	db := mymysql.GetDB(c)
	var count int64 = 0
	//select count(*) from relation where user_id = user_id
	err := db.Table("relation").Where("user_id=? to_user_id=?", userid, touserid).Count(&count).Error
	return count != 0, err
}
