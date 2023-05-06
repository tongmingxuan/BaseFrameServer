// Package Controller /*
package Controller

import (
	"BaseFrameServer/app/Service/AdminService"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	BaseController
}

func (controller AdminController) Login(c *gin.Context) {

	param := controller.FilterData(controller.Params(c), map[string]string{
		"admin_name":     "admin_name:不存在",
		"admin_password": "admin_password:不存在",
	})

	defer func() {
		if r := recover(); r != nil {
			c.String(200, controller.JsonError("异常", nil, 500))
			return
		}
	}()

	result := AdminService.AdminService{}.Login(param["admin_name"].(string), param["admin_password"].(string))

	if result.Code == 200 {
		c.String(200, controller.JsonSuccess(result.Message, map[string]interface{}{
			"token": result.Data,
		}))
	} else {
		c.String(200, controller.JsonError(result.Message, nil, result.Code))
	}
}

func (controller AdminController) GetUserInfo(c *gin.Context) {
	service := AdminService.AdminService{}

	result := service.TokenGetAdmin(c)

	if result.Code != 200 {
		c.String(200, controller.JsonError(result.Message, nil, result.Code))
		return
	}

	c.String(200, controller.JsonSuccess(result.Message, map[string]interface{}{
		"result": result.Data,
	}))
}
