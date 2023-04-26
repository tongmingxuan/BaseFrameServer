// Package bin /*
package bin

import (
	"BaseFrameServer/config"
	"BaseFrameServer/global"
	"github.com/tongmingxuan/tmx-server/tmxServer"
)

// ApplicationStart
// @Description: 应用初始加载
func ApplicationStart() {
	frameList := []tmxServer.BaseFrame{
		&tmxServer.Logger{},
		&tmxServer.DatabasePool{},
		&tmxServer.RedisPool{},
		&tmxServer.Context{},
	}

	configList := []tmxServer.DefaultConfig{
		config.DatabaseConfig{},
		config.RedisConfig{},
		config.ServerConfig{},
	}

	global.Container = tmxServer.ApplicationStart(frameList, configList)
}

func LoadTaskProcess() {
	if tmxServer.GetEnv("task_enable_status", "stop") == "open" {
		task := new(tmxServer.Task)
		task.SetConfig(config.TaskConfig())
		global.GetGlobalContainer().Set(task.Key(), task.Handle())
	}

	if tmxServer.GetEnv("process_enable_status", "stop") == "open" {
		process := new(tmxServer.Process)
		process.SetConfig(config.ProcessConfig())
		global.GetGlobalContainer().Set(process.Key(), process.Handle())
	}
}

func ServerStart() {
	server := new(tmxServer.Server)

	server.SetRouter(config.GetRouteList())

	global.GetGlobalContainer().Set(server.Key(), server.Handle())
}
