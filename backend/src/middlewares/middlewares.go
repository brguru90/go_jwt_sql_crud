package middlewares

import (
	"learn_go/src/my_modules"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func FindUserAgentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if os.Getenv("APP_ENV") != "production" {
		// 	log.WithFields(log.Fields{
		// 		"path":           c.FullPath(),
		// 		"User Agent":     c.GetHeader("User-Agent"),
		// 		"Request Method": c.Request.Method,
		// 	}).Infoln("API Request ==>")
		// }
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
		if c.FullPath() != "" && !strings.HasSuffix(c.FullPath(), "/") {
			c.Redirect(http.StatusTemporaryRedirect, c.FullPath()+"/")
		} else {
			c.Next()
		}
	}
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		decoded_token, err, http_status, ok := my_modules.LoginStatus(c)
		if http_status <= 0 || http_status != 200 {
			my_modules.CreateAndSendResponse(c, http_status, "error", err, nil)
			c.Abort()
		} else if ok {
			c.Set("decoded_token", decoded_token)
			c.Next()
		}
	}
}
