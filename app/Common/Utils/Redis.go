// Package Utils /*
package Utils

import (
	"github.com/go-redis/redis"
)

type Redis struct {
}

func (r Redis) GetRedis(connection string) *redis.Client {
	redisPool := ContainerGetRedis()

	instance, ok := redisPool.Pool.Load(connection)

	if ok != true {
		panic("获取redis链接异常")
	}

	return instance.(*redis.Client)
}

func GetRedisConnection(connection string) *redis.Client {
	return Redis{}.GetRedis(connection)
}
