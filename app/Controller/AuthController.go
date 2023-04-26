// Package Controller /*
package Controller

import "github.com/gin-gonic/gin"

type AuthController struct {
	BaseController
}

func (receiver AuthController) Index(c *gin.Context) {
	c.String(200, receiver.JsonSuccess("嗨害嗨", nil))
}
