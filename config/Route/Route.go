// Package Route /*
package Route

import (
	"github.com/gin-gonic/gin"
)

func DefaultHttpRoute(g *gin.Engine) {

	//g.Use(
	//	Middleware.RequestIdMiddleware{}.MiddlewareFunc(),
	//	Middleware.CorsMiddleware{}.MiddlewareFunc(),
	//)

	//g.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello word")
	//})

	//g.GET("/", Controller.IndexController{}.Index)
}
