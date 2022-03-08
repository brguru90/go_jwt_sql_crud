package apis

import (
	"learn_go/src/apis/views"
	"learn_go/src/middleware"

	"github.com/gin-gonic/gin"
)

// only the functions whose initial letter is upper case only those can be exportable from package
func InitApiTest(router *gin.RouterGroup) {

	router.Use(middleware.ApiSpecificMiddleware())
	router.GET("test/:id", test_api)
	router.GET("hello/", views.Hello_api)
	router.GET("sign_up/", views.SignUp)
	router.GET("login/", views.Login)
	router.GET("login_status/", views.LoginStatus)

	{
		protected_router := router.Group("", middleware.ValidateToken())
		protected_router.GET("user/", views.GetUserData)
		protected_router.PUT("user/", views.UpdateUserData)
		protected_router.DELETE("user/", views.Deleteuser)
		protected_router.GET("user/logout/", views.Deleteuser)
	}

}
