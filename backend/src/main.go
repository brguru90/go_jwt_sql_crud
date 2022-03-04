package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"learn_go/src/apis"

	"github.com/gin-gonic/gin"
)

var all_router *gin.Engine

var SERVER_PORT string = "8000"

func main() {
	// all_router = gin.New()
	all_router = gin.Default()
	api_router := all_router.Group("/api")

	api_router.Use(LogRequest)

	api_router.Use(HeaderHandlerFunc).GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "hi")
	})

	api_router.GET("test2/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "the param sent %s", c.Param("id"))
	})

	apis.InitApiTest(api_router.Group("", FindUserAgentMiddleware()))

	if os.Getenv("SERVER_PORT") != "" {
		SERVER_PORT = os.Getenv("SERVER_PORT")
	}

	// log.Fatal(http.ListenAndServe(":8080", all_router))
	bind_to_host := fmt.Sprintf(":%s", SERVER_PORT)
	all_router.Run(bind_to_host)

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

func LogRequest(c *gin.Context) {
	fmt.Println("request ====> ", c.FullPath())

}

func FindUserAgentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.GetHeader("User-Agent"))
		// Before calling handler
		c.Next()
		// After calling handler
	}
}
