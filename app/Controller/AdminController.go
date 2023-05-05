// Package Controller /*
package Controller

import (
	"BaseFrameServer/app/Service"
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
			c.String(200, controller.JsonError("异常", nil))
			return
		}
	}()

	Service.AdminService{}.Login(param["admin_name"].(string), param["admin_password"].(string))

	c.String(200, controller.JsonSuccess("Login:嗨害嗨", nil))
}
