// Package global /*
package global

import "github.com/tongmingxuan/tmx-server/tmxServer"

var (
	Container *tmxServer.Container
)

func GetGlobalContainer() *tmxServer.Container {
	return Container
}
