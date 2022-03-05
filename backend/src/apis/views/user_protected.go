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

type Row struct {
	Column_id          int64       `json:"id"`
	Column_uuid        string      `json:"uuid"`
	Column_email       string      `json:"email"`
	Column_name        string      `json:"name"`
	Column_description string      `json:"description"`
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
			var rowSlice []Row
			for rows.Next() {
				var r Row
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

	c.String(http.StatusOK, "Welcome hello")
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
