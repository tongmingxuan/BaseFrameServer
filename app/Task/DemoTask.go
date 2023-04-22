// Package Task /*
package Task

import (
	"github.com/tongmingxuan/tmx-server/tmxServer"
	"time"
)

type DemoTask struct {
}

func (t DemoTask) Handle() tmxServer.HandleFunctionTypeOfTask {
	return func() {
		println("demoTask运行")

		time.Sleep(5 * time.Second)
	}
}
