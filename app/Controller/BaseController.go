// Package Controller /*
package Controller

import (
	"BaseFrameServer/app/Exception"
	"github.com/gin-gonic/gin"
	"github.com/syyongx/php2go"
	"github.com/tongmingxuan/tmx-server/plugin/pluginList/HelperFunction"
)

// BaseController
// @Description: 基础控制器类
type BaseController struct {
}

// FilterData
// @Description: 参数验证
// @receiver receiver
// @param params
// @param errorMessageMap
// @return map[string]interface{}
func (receiver BaseController) FilterData(params map[string]interface{}, errorMessageMap map[string]string) map[string]interface{} {
	for checkField, nullMessage := range errorMessageMap {
		value, ok := params[checkField]

		if ok == false {
			Exception.ThrowNewCommonException(500, nullMessage, map[string]interface{}{
				"params":          params,
				"errorMessageMap": errorMessageMap,
			})
		}

		if php2go.Empty(value) {
			Exception.ThrowNewCommonException(500, nullMessage, map[string]interface{}{
				"params":          params,
				"errorMessageMap": errorMessageMap,
			})
		}
	}

	return params
}

// Params
// @Description: 获取参数
// @receiver receiver
// @param c
// @return map[string]interface{}
func (receiver BaseController) Params(c *gin.Context) map[string]interface{} {

	params := make(map[string]interface{})

	for field, value := range HelperFunction.GetQueryParams(c) {
		params[field] = value
	}

	for field, value := range HelperFunction.GetPostFormParams(c) {
		params[field] = value
	}

	return params
}

// JsonSuccess
// @Description: http返回成功
// @receiver receiver
// @param message
// @param data
// @return string
func (receiver BaseController) JsonSuccess(message string, data map[string]interface{}) string {
	return HelperFunction.JsonSuccess(message, data)
}

// JsonError
// @Description: http返回失败
// @receiver receiver
// @param message
// @param data
// @return string
func (receiver BaseController) JsonError(message string, data map[string]interface{}, code int) string {
	return HelperFunction.JsonError(message, data, code)
}
