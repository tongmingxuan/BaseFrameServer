// Package bin /*
package bin

import (
	"fmt"
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

	fmt.Println(frameList)
}
