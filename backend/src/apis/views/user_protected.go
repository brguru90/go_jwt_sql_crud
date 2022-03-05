package views

import (
	"context"
	"fmt"
	"net/http"

	"learn_go/src/my_modules"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type UserRow struct {
	Column_id          int64       `json:"id" binding:"required"`
	Column_uuid        string      `json:"uuid" binding:"required"`
	Column_email       string      `json:"email" binding:"required"`
	Column_name        string      `json:"name" binding:"required"`
	Column_description string      `json:"description" binding:"required"`
	Column_createdAt   interface{} `json:"createdAt"`
	Column_updatedAt   interface{} `json:"updatedAt"`
}

func GetUserData(c *gin.Context, db_connection *pgxpool.Pool) {
	var uuid string = c.Query("uuid")
	if uuid != "" {
		var db_query string = fmt.Sprintf(`select * from users where uuid='%s'; `, uuid)
		rows, err := db_connection.Query(context.Background(), db_query)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"query": db_query,
			}).Errorln("QueryRow failed ==>")
			my_modules.CreateAndSendResponse(c, http.StatusOK, "error", "No record found", nil)
			return
		} else {
			defer rows.Close()
			var rowSlice []UserRow
			for rows.Next() {
				var r UserRow
				err := rows.Scan(&r.Column_id, &r.Column_uuid, &r.Column_name, &r.Column_email, &r.Column_description, &r.Column_createdAt, &r.Column_updatedAt)
				if err != nil {
					log.Errorln(fmt.Sprintf("Scan failed: %v\n", err))
				}
				rowSlice = append(rowSlice, r)
			}
			// log.Debugln("type=%T\nresult=%v", rowSlice, rowSlice)

			if err := rows.Err(); err != nil {
				log.Errorln(fmt.Sprintf("Row Err failed: %v\n", err))
			}

			my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Record found", rowSlice)
			return
		}
	} else {
		my_modules.CreateAndSendResponse(c, http.StatusOK, "error", "Didn't got UUID", nil)
		return
	}
}

func UpdateUserData(c *gin.Context, db_connection *pgxpool.Pool) {
	var updateWithData UserRow
	if err := c.ShouldBindJSON(&updateWithData); err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Invalid input data format", nil)
		return
	}

	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Updated successfully", updateWithData)
}

func GetActiveSession(c *gin.Context, db_connection *pgxpool.Pool) {

	c.String(http.StatusOK, "Welcome hello")
}

func Deleteuser(c *gin.Context, db_connection *pgxpool.Pool) {

	c.String(http.StatusOK, "Welcome hello")
}

func Logout(c *gin.Context, db_connection *pgxpool.Pool) {

	c.String(http.StatusOK, "Welcome hello")
}

func BlockSession(c *gin.Context, db_connection *pgxpool.Pool) {

	c.String(http.StatusOK, "Welcome hello")
}
