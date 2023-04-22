// Package Exception /**
package Exception

// ThrowNewCommonException 通用异常
func ThrowNewCommonException(code int, message string, data map[string]interface{}) {
	CommonException{}.Throw(code, message, data)
}
