package myredis

import "github.com/go-redis/redis"

// IsExistKey key是否寸在
func IsExistKey(key string) bool {
	_, err := client.Get(key).Result()
	//缓存了空值
	if err == redis.Nil {
		return true
	}
	//无key
	if err != nil {
		return false
	}
	return true
}

// GetList 获取redis中的链表
func GetList(key string, start, stop int64) ([]string, error) {
	return client.LRange(key, start, stop).Result()
}

// GetVideoList 获取投稿视频列表
func GetVideoList() ([]string, error) {
	return GetList(key, 0, 10)
}

// SetListForVideoList 更新投稿视频列表
const key = "VideoList"

func SetListForVideoList(val string) {
	client.LTrim("VideoList", 0, 8)
	client.LPushX("VideoList", val)
}
func SetVideoListInit() {
	data := [10]string{"0"}
	client.LPush(key, data)
}
