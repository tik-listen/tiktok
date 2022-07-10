package models

import (
	"context"
	"tiktok/base/mymysql"
)

type Favorite struct {
	FavoriteID int64 `db:"favorite_id"`
	UserID     int64 `db:"user_id"`
	VideoID    int64 `db:"video_id"`
}

//插入点赞
func InsertFavorite(ctx context.Context, favorite Favorite) error {
	db := mymysql.GetDB(ctx)
	err := db.Table("favorite").Create(favorite).Error
	return err
}

//取消点赞
func DeleteFavorite(ctx context.Context, favorite Favorite) error {
	db := mymysql.GetDB(ctx)
	err := db.Table("favorite").Where("user_id=? and video_id=?", favorite.UserID, favorite.VideoID).Delete(favorite).Error
	return err
}

//查询用户点赞视频
func FindFavoriteByUserID(ctx context.Context, UserID int64) ([]Favorite, error) {
	db := mymysql.GetDB(ctx)
	var ret []Favorite
	err := db.Table("favorite").Where("user_id=?", UserID).Find(&ret).Error
	return ret, err
}

//用户是否点赞过
func IsFavorite(ctx context.Context, favorite Favorite) bool {
	db := mymysql.GetDB(ctx)
	var count int64
	db.Table("favorite").Where("user_id=? and video_id=?", favorite.UserID, favorite.VideoID).Count(&count)
	return count != 0
}

// 查询视频点赞数
func CountFavoriteWithEvido(ctx context.Context, videoID int64) (int64, error) {
	db := mymysql.GetDB(ctx)
	var count int64
	err := db.Table("favorite").Where("video_id=?", videoID).Count(&count).Error
	return count, err
}

// 查询视频点赞用户列表
func FindFavoriteByVideoID(ctx context.Context, videoID int64) ([]Favorite, error) {
	db := mymysql.GetDB(ctx)
	ret := make([]Favorite, 0)
	err := db.Table("favorite").Where("video_id=?", videoID).Find(&ret).Error
	return ret, err
}

//查询用户点赞说数
func CountFavoriteWithUser(ctx context.Context, userID int64) (int64, error) {
	db := mymysql.GetDB(ctx)
	var count int64
	err := db.Table("favorite").Where("user_id=?", userID).Count(&count).Error
	return count, err
}
