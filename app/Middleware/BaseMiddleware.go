// Package Middleware /**
package Middleware

import "github.com/gin-gonic/gin"

// BaseMiddleware /**
type BaseMiddleware interface {
	MiddlewareFunc() gin.HandlerFunc
}
