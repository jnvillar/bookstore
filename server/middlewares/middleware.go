package middlewares

import "github.com/gin-gonic/gin"

type Middleware interface {
	RegisterMiddleware(router *gin.Engine)
}
