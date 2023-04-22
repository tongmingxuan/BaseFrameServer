// Package config /*
package config

import (
	"BaseFrameServer/config/Route"
	"github.com/tongmingxuan/tmx-server/tmxServer"
)

func GetRouteList() map[string]tmxServer.RouteReg {
	return map[string]tmxServer.RouteReg{
		"default_http": Route.DefaultHttpRoute,
	}
}
