// Package config /*
package config

import (
	"github.com/tongmingxuan/tmx-server/tmxServer"
)

type ServerConfig struct {
}

func (s ServerConfig) GetConfigKey() string {
	return "server"
}

func (s ServerConfig) GetConfigValue() map[string]map[string]string {
	return map[string]map[string]string{
		"default_http": {
			//监听地址
			"listen_address": tmxServer.GetEnv("listen_address", "0.0.0.0"),
			//监听端口
			"listen_port": tmxServer.GetEnv("listen_port", "9501"),
		},
	}
}
