package models

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"tiktok/base/common"
	"tiktok/base/mymysql"
	"tiktok/base/mymysql/tiktokdb"
)

func CheckUserExist(ctx context.Context, username string) (bool, error) {
	return tiktokdb.CheckUserExist(ctx, username)
}

// InsertOneUser 插入一个用户
func InsertOneUser(ctx context.Context, user *tiktokdb.User) error {

	// 密码加密
	user.Password = encryptPassword(user.Password)

	// 调用公共数据库操作库
	err := tiktokdb.InsertOneUser(ctx, user)
	if err == nil {
		return err
	}

	return nil
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(common.Kmd5Secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *tiktokdb.User) (userId int64, err error) {

	// 用一个 context 控制操作声明周期
	ctx, cancel := context.WithTimeout(context.Background(), mymysql.ConnectTimeout)
	defer cancel()

	res, err := tiktokdb.GetOneUser(ctx, user)

	if err != nil {
		if err == sql.ErrNoRows {
			return -1, common.ErrorInvalidPassword
		}
		// 查询数据库失败
		return -1, common.ErrorDBError
	}

	return res.UserID, nil
}
