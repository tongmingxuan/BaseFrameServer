// Package Utils /*
package Utils

import (
	"github.com/tongmingxuan/tmx-server/tmxServer"
)

// LoggerInfo
//
//	@Description: 日志方法
//	@param message 日志内容
//	@param tag 日志前缀|标识
//	@param data 日志参数
func LoggerInfo(message, tag string, data map[string]interface{}) {
	requestId, ok := tmxServer.GetContext().Get("requestId")

	if ok {
		data["request_id"] = requestId
	}

	ContainerGetLogger().LoggerInfo(message, tag, data)
}
