package logic

import (
	"context"
	"strconv"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/base/jwt"
	"tiktok/base/mymysql/tiktokdb"
	"tiktok/base/snowflake"
	"tiktok/service/favoritesrv/models"

	"github.com/gin-gonic/gin"
)

//点赞和取消赞操作 注册业务操作
func DealLikeAction(ctx context.Context, p *io.LikeActionReq) (*io.Response, error) {
	jwt1, _ := jwt.ParseToken(p.Token)
	userID := jwt1.UserID
	favorite := models.Favorite{UserID: userID, VideoID: p.VideoID}
	exites := models.IsFavorite(ctx, favorite)
	//点赞并且不存在
	if p.ActionType == common.Add && !exites {
		favoriteID := snowflake.GenID()
		favorite.FavoriteID = favoriteID
		err := models.InsertFavorite(ctx, favorite)
		if err != nil {
			return &io.Response{StatusCode: common.CodeInvalidParam, StatusMsg: err.Error()}, nil
		}
		return &io.Response{StatusCode: common.CodeSuccess, StatusMsg: "success"}, err
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

func GetFavoriteList(ctx *gin.Context, p *io.UserInfoReq) (*io.FavoriteListResp, error) {
	favoites, err := models.FindFavoriteByUserID(ctx, p.UserID)
	if err != nil {
		return &io.FavoriteListResp{}, err
	}
	ret := new(io.FavoriteListResp)
	ret.StatusCode = common.CodeSuccess
	ret.StatusMsg = common.CodeSuccess.Msg()
	ret.VideoList = make([]io.VideoRes, 0, len(favoites))
	for _, favoite := range favoites {
		temp, _ := tiktokdb.GetVideoListWithVideoId(ctx, favoite.VideoID)
		res := io.VideoRes{}
		res.Id = temp.VideoId
		res.User, _ = tiktokdb.GetOneUserWithId(ctx, temp.UserId)
		res.PlayUrl = "http://82.157.141.199/" + strconv.FormatInt(temp.VideoId, 10) + ".mp4"
		res.FavoriteCount = temp.FavoriteCount
		res.CommentCount = temp.CommentCount
		res.IsFavorite = temp.IsFavorite
		res.Name = temp.Name
		ret.VideoList = append(ret.VideoList, res)
	}
	return ret, nil
}
