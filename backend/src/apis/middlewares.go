package apis

import (
	"log"

	"github.com/gin-gonic/gin"
)

func ApiSpecificMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("ApiSpecificMiddleware ===>", c.Request.URL.Path)
		// Before calling handler
		c.Next()
		// After calling handler
	}
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("ApiSpecificMiddleware ===>", c.Request.URL.Path)
		// Before calling handler
		c.Next()
		// After calling handler
	}
}
