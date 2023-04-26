// Package BaseFrameServer /*
package main

import (
	"BaseFrameServer/app/Common/Utils"
	"BaseFrameServer/bin"
	"fmt"
	"time"
)

const Father string = `
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

func main() {
	bin.ApplicationStart()

	bin.LoadTaskProcess()

	bin.ServerStart()

	fmt.Println(Father)

	Utils.LoggerInfo("服务启动成功", "", nil)

	time.Sleep(1000 * time.Second)
}
