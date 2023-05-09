package middlewares

import (
	"github.com/gin-gonic/gin"
)

var logSkip = []string{
	"/healthz",
	"/healthz/ready",
	"/metrics",
}

func NewLogger() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: logSkip,
	})
}
