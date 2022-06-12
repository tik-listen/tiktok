package logic

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
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

	// 调用登录操作
	if userId, err = models.Login(p.Username, p.Password); err != nil {
		return -1, "", err
	}

	// 生成JWT
	token, err = jwt.GenToken(userId, p.Username)
	if err != nil {
		return
	}

	return
}

// GetUserInfo 获取用户信息
//func GetUserInfo(ctx context.Context, p *io.UserInfoReq) (resp *io.UserInfoResp, err error) {
//	// 解析 token
//	//claim, err := jwt.ParseToken(p.Token)
//	// 1.判断用户存不存在
//	//flag, err := models.CheckUserExist(ctx, claim.Username)
//	//if err != nil {
//	//	return nil, common.ErrorMysqlDbErr
//	//}
//	//if !flag {
//	//	return nil, common.ErrorUserNotExist
//	//}
//	//
//	resp.ID = 1233
//	resp.FollowerCount = 0
//	resp.FollowCount = 0
//	resp.Name = "claim.Username"
//	resp.IsFollow = false
//	return resp, nil
//}

// GetUserInfo 查询用户信息
func GetUserInfo(ctx *gin.Context, p *io.UserInfoReq, claim *jwt.MyClaims) (*io.UserInfoResp, error) {

	resp := new(io.UserInfoResp)
	userInfo := &tiktokdb.User{
		UserID: p.UserID,
	}
	resp.ID = p.UserID
	// 查询用户信息
	user, err := models.FindOneUser(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	resp.Name = user.Username

	fmt.Println(user)
	fmt.Println(resp)
	// 获取用户的粉丝数
	fansCount, err := models.CountUserFans(ctx, p.UserID)
	if err != nil {
		return nil, err
	}
	resp.FollowerCount = fansCount

	// 获取用户关注数
	followCount, err := models.CountUserStar(ctx, p.UserID)
	if err != nil {
		return nil, err
	}
	resp.FollowCount = followCount

	// 获取是否已经关注
	resp.IsFollow, err = models.IsFans(ctx, claim.UserID, p.UserID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
