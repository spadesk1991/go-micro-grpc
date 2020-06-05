package middleware

import "github.com/gin-gonic/gin"

func AddGRPCServer(s ...interface{}) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("prodGRPCService", s)
		c.Next()
	}
}
