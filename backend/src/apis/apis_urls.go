package apis

import (
	"learn_go/src/apis/api_modules"
	"learn_go/src/apis/views"
	"learn_go/src/middlewares"
	"learn_go/src/my_modules"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const one_sec = 1000000000

var CACHE_TTL, _ = strconv.ParseInt(os.Getenv("RESPONSE_CACHE_TTL_IN_SECS"), 10, 64)
var CACHE_TTL_DURATION = time.Duration(one_sec * CACHE_TTL)

// only the functions whose initial letter is upper case only those can be exportable from package
func InitApiTest(router *gin.RouterGroup) {

	router.Use(middlewares.ApiSpecificMiddleware())

	router.GET("test/:id", test_api)
	router.Any("hello/", views.Hello_api)
	router.Any("hello/:page", views.Hello_api)
	router.Any("hello/:page/:limit", views.Hello_api)
	router.POST("sign_up/", views.SignUp)
	router.POST("login/", views.Login)
	router.GET("login_status/", views.LoginStatus)

	{
		protected_router := router.Group("", middlewares.ValidateToken())
		protected_router.GET("user/", my_modules.GetCachedResponse(views.GetUserData, "users", CACHE_TTL_DURATION, api_modules.ForUserPagination))
		protected_router.PUT("user/", views.UpdateUserData)
		protected_router.DELETE("user/", views.Deleteuser)
		protected_router.GET("user/active_sessions/", views.GetActiveSession)
		protected_router.POST("user/block_token/", views.BlockSession)
		protected_router.GET("user/logout/", views.Logout)
	}

}
