package main

import (
	"fmt"
	"net/http"
	"os"

	"learn_go/src/apis"
	"learn_go/src/database"
	"learn_go/src/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var SERVER_PORT string = "8000"

func main() {

	// https://github.com/gin-gonic/gin

	// all_router = gin.New()
	var all_router *gin.Engine = gin.Default()
	all_router.Use(static.Serve("/", static.LocalFile("./src/static", true)))

	db_connection := database.ConnectPostgres()

	{
		// just grouping, to make it more readable
		api_router := all_router.Group("/api")

		all_router.Use(cors.Default())
		api_router.Use(middleware.FindUserAgentMiddleware()) // an example for global middleware
		api_router.Use(middleware.HeaderHandlerFunc).GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "hi")
		})
		apis.InitApiTest(api_router, db_connection) // more apis imported
	}

	if os.Getenv("SERVER_PORT") != "" {
		SERVER_PORT = os.Getenv("SERVER_PORT")
	}
	// log.Fatal(http.ListenAndServe(":8080", all_router))
	bind_to_host := fmt.Sprintf(":%s", SERVER_PORT) //formatted host string
	all_router.Run(bind_to_host)
}
