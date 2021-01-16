package tool

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
)

type RedisStore struct {
	client *redis.Client
}

var RediStore RedisStore

// 连接、初始化 Redis
func InitRedisStore() *RedisStore {

	config := GetConfig().RedisConfig

	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.Db,
	})

	RediStore = RedisStore{client: client}

	base64Captcha.SetCustomStore(&RediStore)

	return &RediStore

}

// set
func (rs *RedisStore) Set(id string, value string) {
	// time.Minute*10 在程序中缓存10分钟
	err := rs.client.Set(id, value, time.Minute*10).Err()
	if err != nil {
		fmt.Println(err)
	}
}

// get
func (rs *RedisStore) Get(id string, clear bool) string {

	val, err := rs.client.Get(id).Result()

	if err != nil {
		fmt.Println(err)
		return ""
	}

	if clear {
		err := rs.client.Del(id).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}

	return val

}
