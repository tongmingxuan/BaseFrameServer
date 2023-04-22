// Package config /*
package config

import (
	"github.com/tongmingxuan/tmx-server/tmxServer"
)

type DatabaseConfig struct {
}

func (c DatabaseConfig) GetConfigKey() string {
	return "database"
}

func (c DatabaseConfig) GetConfigValue() map[string]map[string]string {
	return map[string]map[string]string{
		"default": {
			"db_host":             tmxServer.GetEnv("db_host", "192.168.2.100"),
			"db_database":         tmxServer.GetEnv("db_database", "base_server"),
			"db_port":             tmxServer.GetEnv("db_port", "6379"),
			"db_username":         tmxServer.GetEnv("db_username", "base_server"),
			"db_password":         tmxServer.GetEnv("db_password", "PthDcbhfJr3E7zBA"),
			"prefix":              tmxServer.GetEnv("prefix", "tmx_"),
			"max_connection":      tmxServer.GetEnv("max_connection", "20"),
			"max_open_connection": tmxServer.GetEnv("max_open_connection", "5"),
		},
	}
}
