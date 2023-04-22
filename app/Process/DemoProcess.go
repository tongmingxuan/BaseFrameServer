// Package Process /*
package Process

import (
	"fmt"
	"github.com/tongmingxuan/tmx-server/tmxServer"
)

type DemoProcess struct {
}

func (d DemoProcess) ProcessHandle(entry tmxServer.ProcessEntry) {
	fmt.Println("processDemo运行", entry)

}
