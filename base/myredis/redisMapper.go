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
