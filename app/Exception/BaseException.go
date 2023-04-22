// Package Exception /**
package Exception

import (
	"reflect"
)

type BaseException interface {
	Throw(code int, message string, data map[string]interface{})
	Handle()
	GetMessage() string
	GetTrace() []string
}

// JudgeTypeOfBaseException 判断是否是异常类型
func JudgeTypeOfBaseException(instance interface{}) bool {
	return reflect.TypeOf(instance).Implements(reflect.TypeOf((*BaseException)(nil)).Elem())
}

// RunHandle 判断是否是异常类型 并 执行对应handle方法
func RunHandle(err interface{}) bool {
	if JudgeTypeOfBaseException(err) == true {
		infoFunc := reflect.ValueOf(err).MethodByName("Handle")
		infoFunc.Call([]reflect.Value{})
		return true
	}

	return false
}
