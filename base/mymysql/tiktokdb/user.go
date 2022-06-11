package tiktokdb

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"tiktok/base/mymysql"
)

const secret = "先暂时写死"

// User 用户
type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"password"`
	Password string `db:"password"`
}

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(ctx context.Context, username string) (bool, error) {

	// 获取数据库连接
	db := mymysql.GetDB(ctx)

	var count int64
	if result := db.Table("users").Where("username = ?", username).Count(&count); result.Error != nil {
		return true, result.Error
	}
	if count > 0 {
		return true, nil
	}

	return false, nil
}

// InsertOneUser 向数据库中插入一条新的用户记录
func InsertOneUser(ctx context.Context, user *User) (err error) {

	// 获取数据库连接
	db := mymysql.GetDB(ctx)

	// 对密码进行加密
	user.Password = encryptPassword(user.Password)

	// 用户注册信息入库
	db.Table("users").Create(user)

	return
}

// GetOneUser 数据库中查询一条用户记录
func GetOneUser(ctx context.Context, user *User) (result User, err error) {

	// 获取数据库连接
	db := mymysql.GetDB(ctx)

	// 对密码进行加密
	user.Password = encryptPassword(user.Password)

	// 查询相关记录
	db.Where(user).Find(&result)

	return
}
func GetOneUserWithId(ctx context.Context, id int64) (User, error) {

	// 获取数据库连接
	db := mymysql.GetDB(ctx)
	var res User
	err := db.Table("user").Where("user_id = ", id).Find(&res)
	if err != nil {
		return res, errors.New("MySQL ERR")
	}
	return res, nil
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
