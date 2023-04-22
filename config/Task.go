// Package config /*
package config

import (
	"BaseFrameServer/app/Task"
	"github.com/tongmingxuan/tmx-server/tmxServer"
)

func TaskConfig() []tmxServer.TaskFrameConfig {
	return []tmxServer.TaskFrameConfig{
		tmxServer.TaskFrameConfig{}.SetName("demo_task").SetMemo("定时任务demo").
			SetRule("@every 5s").SetHandleFunc(Task.DemoTask{}),
	}
}
