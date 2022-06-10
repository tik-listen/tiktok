package logic

import (
	"context"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/base/jwt"
	"tiktok/base/mymysql/tiktokdb"
	"tiktok/base/snowflake"
	"tiktok/service/usersrv/models"
)

// RegisterHandler 注册业务操作
func RegisterHandler(ctx context.Context, p *io.ParamRegister) error {

	// 1.判断用户存不存在
	flag, err := models.CheckUserExist(ctx, p.Username)
	if err != nil {
		return common.ErrorMysqlDbErr
	}
	if flag {
		return common.ErrorUserExist
	}

	// 2.生成 user_id
	userID := snowflake.GenID()

	// 3.构造一个 User 实例
	user := &tiktokdb.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 3.保存进数据库
	return models.InsertOneUser(ctx, user)
}

func Login(p *io.ParamLogin) (userId int64, token string, err error) {

	// 创建要插入的结构
	user := &tiktokdb.User{
		Username: p.Username,
		Password: p.Password,
	}

	// 调用插入操作
	if userId, err = models.Login(user); err != nil {
		return -1, "", err
	}

	// 生成JWT
	token, err = jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}

	return
}
