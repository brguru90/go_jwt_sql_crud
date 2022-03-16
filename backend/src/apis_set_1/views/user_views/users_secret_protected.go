package user_views

import (
	"context"
	"fmt"
	"learn_go/src/database"
	"learn_go/src/my_modules"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const api_secret = "1234"



// @BasePath /api
// @Summary InvalidateUsercache
// @Schemes
// @Description will be used in postgres trigger to delete redis cache
// @Tags Delete user cache
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Param secret header string true "trigger secret"
// @Success 200 {object} my_modules.ResponseFormat
// @Failure 400 {object} my_modules.ResponseFormat
// @Failure 403 {object} my_modules.ResponseFormat
// @Failure 500 {object} my_modules.ResponseFormat
// @Router /del_user_cache/{id} [get]
func InvalidateUsercache(c *gin.Context) {
	if c.GetHeader("secret") != api_secret {
		my_modules.CreateAndSendResponse(c, http.StatusForbidden, "error", "Invalid secret", nil)
		return
	}
	ctx := context.Background()
	

	db_connection := database.POSTGRES_DB_CONNECTION
	db_query := `SELECT uuid FROM users WHERE id=$1`
	rows, err := db_connection.Query(ctx, db_query, c.Param("id"))
	if err != nil {
		if err != context.Canceled {
			log.WithFields(log.Fields{
				"error": err,
				"query": db_query,
			}).Errorln("QueryRow failed ==>")
		}
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "No record found", nil)
		return
	} else {
		defer rows.Close() //importent to prevent connection leak
		for rows.Next() {
			var uuid string
			err := rows.Scan(&uuid)
			if err != nil {
				log.Errorln(fmt.Sprintf("Scan failed: %v\n", err))
				continue
			}
			my_modules.InvalidateCache(uuid)
		}
		if err := rows.Err(); err != nil {
			if err != context.Canceled {
				log.Errorln(fmt.Sprintf("Row Err in rows.Next/rows.Scan failed: %v\n", err))
			}
			my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in retriving user data", nil)
			return
		}
	}

}
