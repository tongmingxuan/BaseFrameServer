// Package Route /*
package Route

import (
	"BaseFrameServer/app/Controller"
	"BaseFrameServer/app/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultHttpRoute(g *gin.Engine) {

	g.Use(
		Middleware.RequestIdMiddleware{}.MiddlewareFunc(),
		Middleware.CorsMiddleware{}.MiddlewareFunc(),
	)

	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})

	g.GET("/auth", Controller.AuthController{}.Index)
}
