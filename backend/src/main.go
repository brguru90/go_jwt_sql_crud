package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"learn_go/src/apis"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

var all_router *gin.Engine

var SERVER_PORT string = "8000"

// https://github.com/gin-gonic/gin

var db_connection *pgxpool.Pool

func main() {
	// all_router = gin.New()
	all_router = gin.Default()
	all_router.Use(static.Serve("/", static.LocalFile("./src/static", true)))

	db_connection = connect_postgres()

	{
		// just grouping, to make it more readable
		api_router := all_router.Group("/api")
		all_router.Use(cors.Default())

		// an example for global middleware
		api_router.Use(FindUserAgentMiddleware())

		api_router.Use(HeaderHandlerFunc).GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "hi")
		})

		// more apis imported
		apis.InitApiTest(api_router, db_connection)
	}

	if os.Getenv("SERVER_PORT") != "" {
		SERVER_PORT = os.Getenv("SERVER_PORT")
	}

	// log.Fatal(http.ListenAndServe(":8080", all_router))
	bind_to_host := fmt.Sprintf(":%s", SERVER_PORT) //formatted host string
	all_router.Run(bind_to_host)

}

func connect_postgres() *pgxpool.Pool {

	var DB_USER string = os.Getenv("DB_USER")
	var DB_PASSWORD string = os.Getenv("DB_PASSWORD")
	var DB_HOST string = os.Getenv("DB_HOST")
	var DATABASE string = os.Getenv("DATABASE")
	var DB_PORT string = os.Getenv("DB_PORT")

	var DB_URL string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DATABASE)

	dbpool, err := pgxpool.Connect(context.Background(), DB_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		fmt.Println("DB_URL ==> ", DB_URL)
		os.Exit(1)
	}
	return dbpool
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
