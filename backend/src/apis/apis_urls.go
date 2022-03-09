package apis

import (
	"learn_go/src/apis/views"
	"learn_go/src/middlewares"

	"github.com/gin-gonic/gin"
)

// only the functions whose initial letter is upper case only those can be exportable from package
func InitApiTest(router *gin.RouterGroup) {

	router.Use(middlewares.ApiSpecificMiddleware())

	router.GET("test/:id", test_api)
	router.GET("hello/", views.Hello_api)
	router.POST("sign_up/", views.SignUp)
	router.POST("login/", views.Login)
	router.GET("login_status/", views.LoginStatus)

	{
		protected_router := router.Group("", middlewares.ValidateToken())
		protected_router.GET("user/", views.GetUserData)
		protected_router.PUT("user/", views.UpdateUserData)
		protected_router.DELETE("user/", views.Deleteuser)
		protected_router.GET("user/active_sessions/", views.GetActiveSession)
		protected_router.GET("user/block_token/", views.BlockSession)
		protected_router.GET("user/logout/", views.Logout)
	}

}
