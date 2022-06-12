package logic

import (
	"context"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/base/jwt"
	"tiktok/service/favoritesrv/models"
)

//点赞和取消赞操作 注册业务操作
func DealLikeAction(ctx context.Context, p *io.LikeActionReq) (*io.Response, error) {
	jwt, _ := jwt.ParseToken(p.Token)
	userID := jwt.UserID
	favorite := models.Favorite{UserID: userID, VideoID: p.VideoID}
	exites := models.IsFavorite(ctx, favorite)
	//点赞并且不存在
	if p.ActionType == common.Add && !exites {

		err := models.InsertFavorite(ctx, favorite)
		if err != nil {
			return &io.Response{StatusCode: common.CodeSuccess, StatusMsg: "success"}, err
		}
		return &io.Response{StatusCode: common.CodeInvalidParam, StatusMsg: err.Error()}, nil
	} else if p.ActionType == common.Cancle && exites {
		//取消点赞且存在
		err := models.DeleteFavorite(ctx, favorite)
		if err != nil {
			return &io.Response{StatusCode: common.CodeInvalidParam, StatusMsg: err.Error()}, err
		}
		return &io.Response{StatusCode: common.CodeSuccess, StatusMsg: "success"}, nil
	} else {
		return &io.Response{StatusCode: common.CodeInvalidParam, StatusMsg: "参数错误"}, common.ErrorInvalid
	}
}

func GetFavoriteList(ctx context.Context, p *io.UserInfoReq) (*io.FavoriteListResp, error) {
	favoites, err := models.FindFavoriteByUserID(ctx, p.UserID)
	if err != nil {
		return &io.FavoriteListResp{}, err
	}
	ret := new(io.FavoriteListResp)
	ret.StatusCode = common.CodeSuccess
	ret.StatusMsg = common.CodeSuccess.Msg()
	for _, favoite := range favoites {
		temp, _ := models.GetVideoWithVideoId(ctx, favoite.VideoID)
		ret.VideoList = append(ret.VideoList, temp)
	}
	return ret, nil
}
