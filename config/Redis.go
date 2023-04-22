// Package config /*
package config

import (
	"github.com/tongmingxuan/tmx-server/tmxServer"
)

type RedisConfig struct {
}

func (r RedisConfig) GetConfigKey() string {
	return "redis"
}

func (r RedisConfig) GetConfigValue() map[string]map[string]string {
	return map[string]map[string]string{
		"default": {
			"redis_host": tmxServer.GetEnv("redis_host", "192.168.2.100"),
			"redis_auth": tmxServer.GetEnv("redis_auth", "123123"),
			"redis_port": tmxServer.GetEnv("redis_port", "6379"),
			"redis_db":   tmxServer.GetEnv("redis_db", "4"),
			//最大连接数
			"redis_max_connection": tmxServer.GetEnv("redis_max_connection", "20"),
			//最小空闲连接
			"redis_min_open_connection": tmxServer.GetEnv("redis_min_open_connection", "5"),
		},
	}
}
