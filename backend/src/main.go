package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"learn_go/src/apis"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var all_router *gin.Engine

var SERVER_PORT string = "8000"

func main() {
	// all_router = gin.New()
	all_router = gin.Default()
	all_router.Use(static.Serve("/", static.LocalFile("./src/static", true)))

	{
		// just grouping, to make it more readable
		api_router := all_router.Group("/api")

		// an example for global middleware
		api_router.Use(FindUserAgentMiddleware())

		api_router.Use(HeaderHandlerFunc).GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "hi")
		})

		// more apis imported
		apis.InitApiTest(api_router)
	}

	if os.Getenv("SERVER_PORT") != "" {
		SERVER_PORT = os.Getenv("SERVER_PORT")
	}

	// log.Fatal(http.ListenAndServe(":8080", all_router))
	bind_to_host := fmt.Sprintf(":%s", SERVER_PORT) //formatted host string
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

func FindUserAgentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.FullPath(), " | User Agent Logger ===>", c.GetHeader("User-Agent"))
		// Before calling handler
		c.Next()
		// After calling handler
	}
}
