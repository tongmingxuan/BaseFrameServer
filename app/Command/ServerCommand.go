// Package Command /*
package Command

import (
	"BaseFrameServer/bin"
	"fmt"
	"github.com/urfave/cli"
)

const Father string = `
	                   嗨害嗨,大爹保佑
	                    _ooOoo_
	                   o8888888o
	                   88" . "88
	                   (| -_- |)
	                    O\ = /O
	                ____/'---'\____
	              .   ' \\| |// '.
	               / \\||| : |||// \
	             / _||||| -:- |||||- \
	               | | \\\ - /// | |
	             | \_| ''\---/'' | |
	              \ .-\__  - ___/-. /
	           ___ . .' /--.--\  . . __
	        ."" '<  .___\_<|>_/___.' >'"".
	       | | :  - \'.;'\ _ /';.'/ -   : | |
	         \ \  -. \_ __\ /__ _/ .-  / /
	 ====== -.____ -.___\_____/___.- ____.-'======

  ____   _   _   ____    _____    ___    __  __   _____   ____  
 / ___| | | | | / ___|  |_   _|  / _ \  |  \/  | | ____| |  _ \ 
| |     | | | | \___ \    | |   | | | | | |\/| | |  _|   | |_) |
| |___  | |_| |  ___) |   | |   | |_| | | |  | | | |___  |  _ < 
 \____|  \___/  |____/    |_|    \___/  |_|  |_| |_____| |_| \_\
                                                                                       
`

// ServerCommand
// @Description: 主服务运行
type ServerCommand struct {
}

func (command ServerCommand) Command() cli.Command {
	return cli.Command{
		Name:  "server",
		Usage: "run",
		Subcommands: []cli.Command{
			{
				Name:   "start",
				Usage:  "开启运行server服务",
				Action: command.Run,
			},
		},
	}
}

func (command ServerCommand) Run(c *cli.Context) error {
	bin.LoadTaskProcess()

	bin.ServerStart()

	//输出大爹
	fmt.Println(Father, "")

	select {}
}
