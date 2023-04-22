// Package config /*
package config

import (
	"BaseFrameServer/app/Process"
	"github.com/tongmingxuan/tmx-server/tmxServer"
)

func ProcessConfig() []tmxServer.ProcessConfig {
	return []tmxServer.ProcessConfig{
		{
			ProcessName:   "demo",
			ProcessNumber: 3,
			ProcessFunc:   Process.DemoProcess{},
		},
	}
}
