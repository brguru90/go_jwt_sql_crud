package apis_set_1

import (
	"learn_go/src/apis_set_1/api_modules"
	"learn_go/src/apis_set_1/views"
	"learn_go/src/apis_set_1/views/user_views"
	"learn_go/src/configs"
	"learn_go/src/middlewares"
	"learn_go/src/my_modules"
	"time"

	"github.com/gin-gonic/gin"
)

const one_sec = 1000000000

// only the functions whose initial letter is upper case only those can be exportable from package
func InitApiTest(router *gin.RouterGroup) {
	var CACHE_TTL_DURATION = time.Duration(one_sec * configs.EnvConfigs.RESPONSE_CACHE_TTL_IN_SECS)

	router.Use(middlewares.ApiSpecificMiddleware())

	router.GET("test/:id", test_api)
	router.Any("hello/", views.Hello_api)
	router.Any("hello/:page", views.Hello_api)
	router.Any("hello/:page/:limit", views.Hello_api)
	router.POST("sign_up/", user_views.SignUp)
	router.POST("login/", user_views.Login)
	router.GET("login_status/", user_views.LoginStatus)
	router.GET("del_user_cache/:id", user_views.InvalidateUsercache)

	{
		protected_router := router.Group("", middlewares.ValidateToken())
		protected_router.GET("user/", my_modules.GetCachedResponse(user_views.GetUserData, "users", CACHE_TTL_DURATION, api_modules.ForUserPagination))
		protected_router.PUT("user/", user_views.UpdateUserData)
		protected_router.DELETE("user/", user_views.Deleteuser)
		protected_router.GET("user/active_sessions/", user_views.GetActiveSession)
		protected_router.POST("user/block_token/", user_views.BlockSession)
		protected_router.GET("user/logout/", user_views.Logout)
	}

}
