package logic

import (
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/base/jwt"
	"tiktok/base/snowflake"
	"tiktok/service/usersrv/models"

	"github.com/gin-gonic/gin"
)

func DealRelationAction(c *gin.Context, relation *io.ParamRealation) (*io.Response, error) {
	if relation.ActionType == common.Add {
		//生成一个雪花id
		relationID := snowflake.GenID()
		r := models.Relation{RelationID: relationID, UserID: relation.UserID, ToUserID: relation.ToUserID}
		status := models.InsertRelation(c, r)
		return &io.Response{StatusCode: common.CodeSuccess, StatusMsg: "关注成功"}, status
	} else if relation.ActionType == common.Cancle {
		r := models.Relation{UserID: relation.UserID, ToUserID: relation.ToUserID}
		status := models.DeleteRelation(c, r)
		return &io.Response{StatusCode: common.CodeSuccess, StatusMsg: "取消关注成功"}, status
	} else {
		return nil, common.ErrorInvalid
	}
}
func FindFollweList(c *gin.Context, r *io.UserInfoReq) (*io.RelationResponse, error) {
	var relations []models.Relation
	relations, err := models.FindUserStar(c, r.UserID)
	ret := new(io.RelationResponse)
	if err != nil {
		return ret, err
	}
	ret.Response.StatusCode = common.CodeSuccess
	ret.Response.StatusMsg = "success"
	ret.UserList = make([]io.UserInfoResp, len(relations))
	//找到关注列表信息，所以应该是ToUserInfo
	myclas, err := jwt.ParseToken(r.Token)
	if err != nil {
		return &io.RelationResponse{}, err
	}
	for _, relation := range relations {
		user := relation.ToUserID
		userinfo, err := GetUserInfo(c, &io.UserInfoReq{UserID: user, Token: r.Token}, myclas)
		if err != nil {
			return ret, err
		}
		ret.UserList = append(ret.UserList, *userinfo)
	}
	return ret, nil

}

func FindFollwerList(c *gin.Context, r *io.UserInfoReq) (*io.RelationResponse, error) {
	relations, err := models.FindUserFans(c, r.UserID)
	ret := new(io.RelationResponse)
	if err != nil {
		return ret, err
	}
	ret.Response.StatusCode = common.CodeSuccess
	ret.Response.StatusMsg = "success"
	ret.UserList = make([]io.UserInfoResp, len(relations))
	//找到我的粉丝 所以应该是UserID
	myclas, err := jwt.ParseToken(r.Token)
	if err != nil {
		return &io.RelationResponse{}, err
	}
	for _, relation := range relations {
		user := relation.UserID
		userinfo, err := GetUserInfo(c, &io.UserInfoReq{UserID: user, Token: r.Token}, myclas)
		if err != nil {
			return ret, err
		}
		ret.UserList = append(ret.UserList, *userinfo)
	}
	return ret, nil
}
