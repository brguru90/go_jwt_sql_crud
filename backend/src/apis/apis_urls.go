package apis

import (
	"learn_go/src/apis/views"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

// only the functions whose initial letter is upper case only those can be exportable from package
func InitApiTest(router *gin.RouterGroup, db_connection *pgxpool.Pool) {
	router.Use(ApiSpecificMiddleware())
	router.GET("test/:id", test_api)
	router.GET("hello/", views.Hello_api)
}
