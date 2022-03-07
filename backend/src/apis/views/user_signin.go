package views

import (
	"learn_go/src/my_modules"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SignUp(c *gin.Context, db_connection *pgxpool.Pool) {

	var newUserRow NewUserRow
	if err := c.ShouldBindJSON(&newUserRow); err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Invalid input data format", nil)
		return
	}

	my_modules.GenerateAccessToken(
		newUserRow.Column_email,
		"1234",
		map[string]string{
			"email":newUserRow.Column_email,
			"uuid":newUserRow.Column_uuid,
		},
	)

	c.String(http.StatusOK, "Welcome hello")
}

func Login(c *gin.Context, db_connection *pgxpool.Pool) {

	c.String(http.StatusOK, "Welcome hello")
}
