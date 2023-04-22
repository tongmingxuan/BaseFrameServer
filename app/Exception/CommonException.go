package Exception

import (
	"BaseFrameServer/app/Common/Utils"
	"github.com/tongmingxuan/tmx-server/tmxServer"
)

type CommonException struct {
	Code    int
	Message string
	Data    map[string]interface{}
	Trace   []string
}

func (exception CommonException) Throw(code int, message string, data map[string]interface{}) {
	exception.Code = code
	exception.Message = message
	exception.Data = data
	exception.Trace = tmxServer.GetDebugTraceBySlice()

	Utils.LoggerInfo("error:CommonException:Throw:"+message, "error:CommonException", map[string]interface{}{
		"error": exception,
	})

	panic(exception)
}

func (exception CommonException) Handle() {

	Utils.LoggerInfo("error:CommonException:Throw:Handle:"+exception.Message, "error:CommonException", map[string]interface{}{
		"error": exception,
	})
	//todo 自定义是否发送异常消息提醒

	//todo 自定义是否异常信息入队

	//todo 自定义等等等
}

func (exception CommonException) GetMessage() string {
	return exception.Message
}

func (exception CommonException) GetTrace() []string {
	return exception.Trace
}
