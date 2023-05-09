// Package BaseFrameServer /*
package main

import (
	"BaseFrameServer/app/Command"
	"BaseFrameServer/bin"
	"BaseFrameServer/global"
	"github.com/tongmingxuan/tmx-server/tmxServer"
)

func main() {
	bin.ApplicationStart()

	command := new(tmxServer.Command)

	command.SetConfig([]tmxServer.CommandItem{
		Command.ServerCommand{},
		Command.DemoCommand{},
	})

	global.GetGlobalContainer().Set(command.Key(), command.Handle())
}
