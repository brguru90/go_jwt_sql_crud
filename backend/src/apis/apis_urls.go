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
	router.GET("sign_up/", func(c *gin.Context) {
		views.SignUp(c, db_connection)
	})
	router.GET("login/", func(c *gin.Context) {
		views.Login(c, db_connection)
	})

	{
		protected_router := router.Group("", ValidateToken())

		protected_router.GET("user/", func(c *gin.Context) {
			views.GetUserData(c, db_connection)
		})

		protected_router.PUT("user/", func(c *gin.Context) {
			views.UpdateUserData(c, db_connection)
		})

		protected_router.DELETE("user/", func(c *gin.Context) {
			views.Deleteuser(c, db_connection)
		})

		protected_router.GET("user/logout/", func(c *gin.Context) {
			views.Deleteuser(c, db_connection)
		})
	}

}
