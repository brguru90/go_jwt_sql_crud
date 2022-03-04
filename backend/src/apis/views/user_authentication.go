package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SignUp(c *gin.Context, db_connection *pgxpool.Pool) {

	c.String(http.StatusOK, "Welcome hello")
}

func Login(c *gin.Context, db_connection *pgxpool.Pool) {

	c.String(http.StatusOK, "Welcome hello")
}
