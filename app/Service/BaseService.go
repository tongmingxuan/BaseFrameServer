// Package Service /*
package Service

type BaseService struct {
}

type Result struct {
	Code    int
	Message string
	Data    interface{}
}

// ServiceSuccess
// @Description: service成功
// @receiver service
// @param message
// @param code
// @param data
func (service BaseService) ServiceSuccess(message string, data interface{}) Result {
	return Result{
		Code:    200,
		Message: message,
		Data:    data,
	}
}

// ServiceError
// @Description: service异常
// @receiver service
// @param message
// @param data
// @return Result
func (service BaseService) ServiceError(message string, data interface{}, code int) Result {
	return Result{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
