package views

import (
	"context"
	"fmt"
	"net/http"

	"learn_go/src/database"
	"learn_go/src/my_modules"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserData(c *gin.Context) {
	db_connection := database.POSTGRES_DB_CONNECTION
	var uuid string = c.Query("uuid")
	if uuid != "" {
		var db_query string = fmt.Sprintf(`SELECT * FROM users WHERE uuid='%s'; `, uuid)
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

func UpdateUserData(c *gin.Context) {
	db_connection := database.POSTGRES_DB_CONNECTION
	var updateWithData UserRow
	if err := c.ShouldBindJSON(&updateWithData); err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Invalid input data format", nil)
		return
	}

	const sql_stmt string = `UPDATE users SET email=$2,name=$3,description=$4 WHERE uuid=$1`
	res, err := db_connection.Exec(context.Background(), sql_stmt, updateWithData.Column_uuid, updateWithData.Column_email, updateWithData.Column_name, updateWithData.Column_description)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"sql": fmt.Sprintf(`UPDATE users SET email=%s,name=%s,description=%s WHERE uuid=%s`, updateWithData.Column_email, updateWithData.Column_name, updateWithData.Column_description, updateWithData.Column_uuid),
		}).Errorln("Failed to update user data")
		my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Failed to update data", nil)
		return
	}

	rows_updated := res.RowsAffected()

	var response_data = make(map[string]interface{})
	response_data["updated_with_data"] = updateWithData
	response_data["updated_count"] = rows_updated

	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Updated successfully", response_data)
}

func GetActiveSession(c *gin.Context) {
	c.String(http.StatusOK, "Welcome hello")
}

func Deleteuser(c *gin.Context) {
	db_connection := database.POSTGRES_DB_CONNECTION
	var uuid string = c.Query("uuid")

	if uuid == "" {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "UUID of user is not provided", nil)
		return
	}

	const sql_stmt string = `DELETE FROM users WHERE uuid=$1`
	res, err := db_connection.Exec(context.Background(), sql_stmt, uuid)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"sql": fmt.Sprintf(`DELETE FROM users WHERE uuid=%s`, uuid),
		}).Errorln("Failed to delete user data")
		my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Failed to delete user data", nil)
		return
	}

	rows_deleted := res.RowsAffected()

	var response_data = make(map[string]interface{})
	response_data["deleted_user_with_uuid"] = uuid
	response_data["deleted_count"] = rows_deleted

	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Updated successfully", response_data)

}

func Logout(c *gin.Context) {

	c.String(http.StatusOK, "Welcome hello")
}

func BlockSession(c *gin.Context) {

	c.String(http.StatusOK, "Welcome hello")
}
