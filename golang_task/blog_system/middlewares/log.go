package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}

		log.Printf("%s %s %s %s %d %s",
			c.ClientIP(),
			c.Request.Method,
			path,
			c.Request.Proto,
			c.Writer.Status(),
			time.Since(start),
		)
	}
}
