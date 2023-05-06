// Package Middleware requestId中间件
package Middleware

import (
	"BaseFrameServer/app/Common/Utils"
	"BaseFrameServer/app/Exception"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/tongmingxuan/tmx-server/plugin/pluginList/HelperFunction"
	"github.com/tongmingxuan/tmx-server/tmxServer"
	"reflect"
	"runtime"
)

// RequestIdMiddleware
// @Description: http request_id记录中间件
type RequestIdMiddleware struct {
}

// 自定义一个结构体，实现 gin.ResponseWriter interface
type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

// 重写 Write([]byte) (int, error) 方法
func (w responseWriter) Write(b []byte) (int, error) {
	//向一个bytes.buffer中写一份数据来为获取body使用
	w.b.Write(b)
	//完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}

func (m RequestIdMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			tmxServer.GetContext().Release()
		}()

		defer func() {
			if r := recover(); r != nil {
				handlerErr := Exception.RunHandle(r)

				switch r.(type) {
				case string: //手动panic异常
					context.String(200, HelperFunction.JsonError("string.error", map[string]interface{}{
						"error":     r,
						"RunHandle": handlerErr,
					}, 500))
				case runtime.Error: // 运行时错误:
					context.String(200, HelperFunction.JsonError("runtime.error", map[string]interface{}{
						"error":     r,
						"RunHandle": handlerErr,
					}, 500))
				case Exception.BaseException:
					message := reflect.ValueOf(r).FieldByName("Message").String()
					context.String(200, HelperFunction.JsonError("customize.error:"+message, map[string]interface{}{
						"error":     r,
						"RunHandle": handlerErr,
					}, 500))
				default:
					context.String(200, HelperFunction.JsonError("运行异常", map[string]interface{}{
						"error":     r,
						"RunHandle": handlerErr,
					}, 500))
				}
			}
		}()

		requestId := HelperFunction.GetSnowflakeIdByInt64()

		tmxServer.GetContext().Set("requestId", requestId)

		getParams := HelperFunction.GetQueryParams(context)

		postParams := HelperFunction.GetPostFormParams(context)

		Utils.LoggerInfo("RequestIdMiddleware", "接收到请求", map[string]interface{}{
			"message":      "接收到请求",
			"get_params":   getParams,
			"post_params":  postParams,
			"request_id":   requestId,
			"request_path": context.Request.RequestURI,
		})

		writer := responseWriter{
			context.Writer,
			bytes.NewBuffer([]byte{}),
		}

		context.Writer = writer

		context.Next()

		Utils.LoggerInfo("RequestIdMiddleware", "返回响应", map[string]interface{}{
			"message":      "返回响应",
			"get_params":   getParams,
			"post_params":  postParams,
			"request_id":   requestId,
			"request_path": context.Request.RequestURI,
			"res":          writer.b.String(),
		})
	}
}
