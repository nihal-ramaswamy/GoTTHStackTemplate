package htmx_middleware

import "github.com/gin-gonic/gin"

func TextHtmlMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html charset=utf-8")
	}
}
