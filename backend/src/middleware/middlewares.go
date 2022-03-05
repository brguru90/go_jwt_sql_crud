package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func FindUserAgentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Infoln(" | ", c.FullPath(), " | User Agent Logger ===>", c.GetHeader("User-Agent"))
		// Before calling handler
		c.Next()
		// After calling handler
	}
}

func HeaderHandlerFunc(c *gin.Context) {
	h := c.GetHeader("token")
	if h == "1234" {
		c.Header("User", "some user")
	} else if h == "" {
		c.Header("User", "no token received")
		// will abort request to this middleware
		// c.AbortWithStatus(http.StatusOK)
	} else {
		c.Header("User", "invalid token")
	}
}

func ApiSpecificMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// log.Debugln("ApiSpecificMiddleware ===>", c.Request.URL.Path)
		// Before calling handler
		c.Next()
		// After calling handler
	}
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// log.Debugln("ValidateToken,ApiSpecificMiddleware ===>", c.Request.URL.Path)
		// Before calling handler
		c.Next()
		// After calling handler
	}
}
