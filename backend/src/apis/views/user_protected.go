package views

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"learn_go/src/database"
	"learn_go/src/my_modules"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

func GetUserData(c *gin.Context) {
	db_connection := database.POSTGRES_DB_CONNECTION
	payload, ok := my_modules.ExtractTokenPayload(c)
	if !ok {
		return
	}
	var uuid string = payload.UUID

	var _limit int64 = 100
	var _page int64 = 0

	if c.Query("page") != "" {
		_page, _ = strconv.ParseInt(c.Query("page"), 10, 64)
		if c.Query("limit") != "" {
			_limit, _ = strconv.ParseInt(c.Query("limit"), 10, 64)
		}
	}

	if uuid != "" {
		var db_query string
		var rows pgx.Rows
		var err error

		if _page > 0 {
			// this pagination is just implemented to benchmark the api have multiple record
			// for now lets assume admin as current use
			var _offset = _limit * (_page - 1)
			db_query = `SELECT * FROM users ORDER BY id OFFSET $2 LIMIT $1; `
			rows, err = db_connection.Query(c.Request.Context(), db_query, _limit, _offset)
		} else {
			db_query = `SELECT * FROM users WHERE uuid=$1`
			rows, err = db_connection.Query(c.Request.Context(), db_query, uuid)
		}

		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"query": db_query,
			}).Errorln("QueryRow failed ==>")
			my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "No record found", nil)
			return
		} else {
			defer rows.Close() //importent to prevent connection leak
			var rowSlice []my_modules.UserRow
			for rows.Next() {
				var r my_modules.UserRow
				err := rows.Scan(&r.Column_id, &r.Column_uuid, &r.Column_name, &r.Column_email, &r.Column_description, &r.Column_createdAt, &r.Column_updatedAt)
				if err != nil {
					log.Errorln(fmt.Sprintf("Scan failed: %v\n", err))
				}
				rowSlice = append(rowSlice, r)
			}
			// log.Debugln("type=%T\nresult=%v", rowSlice, rowSlice)

			if err := rows.Err(); err != nil {
				log.Errorln(fmt.Sprintf("Row Err in rows.Next/rows.Scan failed: %v\n", err))
				my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in retriving user data", nil)
			}

			if _page > 0 {
				log.Errorln(fmt.Sprintf("GetUserData - Rows count: %v\n", len(rowSlice)))
				my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Record found", map[string]interface{}{
					"users":    rowSlice,
					"cur_page": _page,
				})
				return
			} else {
				my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Record found", rowSlice[0])
				return
			}

		}
	} else {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Didn't got UUID", nil)
		return
	}
}

type NewUserDataFormat struct {
	NewUserData my_modules.NewUserRow `json:"newUserData" binding:"required"`
}

func UpdateUserData(c *gin.Context) {
	db_connection := database.POSTGRES_DB_CONNECTION
	payload, ok := my_modules.ExtractTokenPayload(c)
	if !ok {
		return
	}

	var newUserData NewUserDataFormat

	if err := c.ShouldBindJSON(&newUserData); err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Invalid input data format", nil)
		return
	}
	var updateWithData my_modules.NewUserRow = newUserData.NewUserData

	_time := time.Now()
	updateWithData.Column_uuid = payload.UUID

	const sql_stmt string = `UPDATE users SET email=$2,name=$3,description=$4,"updatedAt"=$5 WHERE uuid=$1`
	res, err := db_connection.Exec(context.Background(), sql_stmt, updateWithData.Column_uuid, updateWithData.Column_email, updateWithData.Column_name, updateWithData.Column_description, _time)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"sql": fmt.Sprintf(`UPDATE users SET email=%s,name=%s,description=%s,updatedAt=%v WHERE uuid=%s`, updateWithData.Column_email, updateWithData.Column_name, updateWithData.Column_description, _time, updateWithData.Column_uuid),
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
	payload, ok := my_modules.ExtractTokenPayload(c)
	if !ok {
		return
	}

	var uuid string = payload.UUID

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
	if rows_deleted > 0 {
		my_modules.DeleteCookie(c, "access_token")
		my_modules.DeleteCookie(c, "user_data")
	}

	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Updated successfully", response_data)

}

func Logout(c *gin.Context) {
	my_modules.DeleteCookie(c, "access_token")
	my_modules.DeleteCookie(c, "user_data")
	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Logged out", nil)
}

func BlockSession(c *gin.Context) {

	c.String(http.StatusOK, "Welcome hello")
}
