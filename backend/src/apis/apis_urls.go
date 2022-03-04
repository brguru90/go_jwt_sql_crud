package apis

import (
	"learn_go/src/apis/views"

	"github.com/gin-gonic/gin"
)

// only the functions whose initial letter is upper case only those can be exportable from package
func InitApiTest(router *gin.RouterGroup) {
	router.Use(ApiSpecificMiddleware())
	router.GET("test/:id", test_api)
	router.GET("hello/", views.Hello_api)
}
