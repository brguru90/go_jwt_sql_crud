package apis

import (
	"github.com/gin-gonic/gin"
)

// only the functions whose initial letter is upper case only those can be exportable from package
func InitApiTest(router *gin.RouterGroup) {
	router.Use(apiSpecificMiddleware())
	router.GET("hello/", hello_api)
}
