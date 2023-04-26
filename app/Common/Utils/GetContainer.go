// Package Utils /*
package Utils

import (
	"BaseFrameServer/global"
	"github.com/tongmingxuan/tmx-server/tmxServer"
	"unsafe"
)

// ContainerGetLogger
//
//	@Description: 获取日志实例
//	@return *Frame.Logger
func ContainerGetLogger() *tmxServer.Logger {
	return (*tmxServer.Logger)(unsafe.Pointer(global.GetGlobalContainer().Get(new(tmxServer.Logger).Key())))
}

// ContainerGetDataBase
//
//	@Description: 获取数据库连接
//	@return *Frame.DatabasePool
func ContainerGetDataBase() *tmxServer.DatabasePool {
	return (*tmxServer.DatabasePool)(unsafe.Pointer(global.GetGlobalContainer().Get(new(tmxServer.DatabasePool).Key())))
}

// ContainerGetRedis
//
//	@Description: 获取redis
//	@return *Frame.RedisPool
func ContainerGetRedis() *tmxServer.RedisPool {
	return (*tmxServer.RedisPool)(unsafe.Pointer(global.GetGlobalContainer().Get(new(tmxServer.RedisPool).Key())))
}

// ContainerGetConfig
//
//	@Description: 获取config
//	@return *Frame.Config
func ContainerGetConfig() *tmxServer.Config {
	return (*tmxServer.Config)(unsafe.Pointer(global.GetGlobalContainer().Get(new(tmxServer.Config).Key())))
}

// ContainerGetServer
// @Description: 获取server
// @return *tmxServer.Config
func ContainerGetServer() *tmxServer.Server {
	return (*tmxServer.Server)(unsafe.Pointer(global.GetGlobalContainer().Get(new(tmxServer.Server).Key())))
}
