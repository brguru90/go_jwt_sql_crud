package user_views

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

// @BasePath /api
// @Summary url to view user data
// @Schemes
// @Description allow people to their user profile data
// @Tags View user data
// @Accept json
// @Produce json
// @Param page query string false "page"
// @Param limit query string false "limit"
// @Success 200 {object} my_modules.ResponseFormat
// @Failure 400 {object} my_modules.ResponseFormat
// @Failure 500 {object} my_modules.ResponseFormat
// @Router /user/ [get]
func GetUserData(c *gin.Context) {
	db_connection := database.POSTGRES_DB_CONNECTION
	payload, ok := my_modules.ExtractTokenPayload(c)
	if !ok {
		return
	}
	var uuid string = payload.Data.UUID

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
			// since its an pagination query sending request context, so query can be terminated if API request is cancelled
			rows, err = db_connection.Query(c.Request.Context(), db_query, _limit, _offset)
		} else {
			db_query = `SELECT * FROM users WHERE uuid=$1`
			rows, err = db_connection.Query(context.Background(), db_query, uuid)
		}

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
			var rowSlice []my_modules.UserRow
			for rows.Next() {
				var r my_modules.UserRow
				err := rows.Scan(&r.Column_id, &r.Column_uuid, &r.Column_name, &r.Column_email, &r.Column_description, &r.Column_createdAt, &r.Column_updatedAt)
				if err != nil {
					log.Errorln(fmt.Sprintf("Scan failed: %v\n", err))
					continue
				}
				rowSlice = append(rowSlice, r)
			}
			// log.Debugln("type=%T\nresult=%v", rowSlice, rowSlice)

			if err := rows.Err(); err != nil {
				if err != context.Canceled {
					log.Errorln(fmt.Sprintf("Row Err in rows.Next/rows.Scan failed: %v\n", err))
				}
				my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in retriving user data", nil)
				return
			}

			if _page > 0 {
				total_users,_:=database.REDIS_DB_CONNECTION.Get(context.Background(),"users_count").Result()
				total_users_int,_:=strconv.ParseInt(total_users,10,64)
				my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Record found", map[string]interface{}{
					"users":    rowSlice,
					"cur_page": _page,
					"total_users":total_users_int,
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
	updateWithData.Column_uuid = payload.Data.UUID

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
	db_connection := database.POSTGRES_DB_CONNECTION
	payload, ok := my_modules.ExtractTokenPayload(c)
	if !ok {
		return
	}
	var uuid string = payload.Data.UUID

	db_query := `SELECT * FROM active_sessions WHERE user_uuid=$1 and token_id!=$2`
	rows, err := db_connection.Query(c.Request.Context(), db_query, uuid, payload.Token_id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"query": db_query,
		}).Errorln("QueryRow failed ==>")
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "No record found", nil)
		return
	} else {
		defer rows.Close() //importent to prevent connection leak
		var rowSlice []my_modules.ActiveSessionsRow = []my_modules.ActiveSessionsRow{}
		for rows.Next() {
			var r my_modules.ActiveSessionsRow
			err := rows.Scan(&r.Column_id, &r.Column_uuid, &r.Column_user_uuid, &r.Column_token_id, &r.Column_ua, &r.Column_ip, &r.Column_exp, &r.Column_status, &r.Column_createdAt, &r.Column_updatedAt)
			if err != nil {
				log.Errorln(fmt.Sprintf("Scan failed: %v\n", err))
			}
			rowSlice = append(rowSlice, r)
		}

		if err := rows.Err(); err != nil {
			log.Errorln(fmt.Sprintf("Row Err in rows.Next/rows.Scan failed: %v\n", err))
			my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in retriving active session", nil)
			return
		}

		my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Record found", rowSlice)
		return
	}
}

func Deleteuser(c *gin.Context) {
	db_connection := database.POSTGRES_DB_CONNECTION
	payload, ok := my_modules.ExtractTokenPayload(c)
	if !ok {
		return
	}

	var uuid string = payload.Data.UUID

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

func updateActiveSession(activeSessionsRow my_modules.ActiveSessionsRow, status string) (int64, error) {
	db_connection := database.POSTGRES_DB_CONNECTION

	var sql_stmt string = `UPDATE active_sessions SET status=$1 WHERE token_id=$2 AND user_uuid=$3 AND exp=$4`
	if status == "blocked" {
		sql_stmt += ` AND status='active'`
	}
	res, err := db_connection.Exec(context.Background(), sql_stmt, status, activeSessionsRow.Column_token_id, activeSessionsRow.Column_user_uuid, activeSessionsRow.Column_exp)

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"sql": fmt.Sprintf(`UPDATE active_sessions SET status=%s WHERE token_id=%s,user_uuid=%s,exp=%d,status="active"`, "blocked", activeSessionsRow.Column_token_id, activeSessionsRow.Column_user_uuid, activeSessionsRow.Column_exp),
		}).Errorln("Failed to update user data")
		return -1, err
	}
	rows_updated := res.RowsAffected()
	return rows_updated, nil
}

func BlockSession(c *gin.Context) {
	redis_db_connection := database.REDIS_DB_CONNECTION
	payload, ok := my_modules.ExtractTokenPayload(c)
	if !ok {
		return
	}

	var activeSessionsRow my_modules.ActiveSessionsRow
	if err := c.ShouldBindJSON(&activeSessionsRow); err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Invalid input data format", nil)
		return
	}
	activeSessionsRow.Column_user_uuid = payload.Data.UUID

	rows_updated, err := updateActiveSession(activeSessionsRow, "blocked")
	if err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Failed to update data", nil)
		return
	}
	if rows_updated <= 0 {
		my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Token doesn't exists/Already blacklisted", nil)
		return
	}
	var exp_sec time.Duration = time.UnixMilli(activeSessionsRow.Column_exp).UTC().Sub(time.Now().UTC())
	r_err := redis_db_connection.SetEX(context.Background(), activeSessionsRow.Column_token_id, activeSessionsRow.Column_user_uuid,
		exp_sec).Err()
	if r_err != nil {
		rows_updated, err := updateActiveSession(activeSessionsRow, "active")
		log.WithFields(log.Fields{
			"redis_err":        r_err,
			"sql_err":          err,
			"sql_rows_updated": rows_updated,
		}).Errorln("Failed to insert data on redis")
		my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Failed to blacklist the session", nil)
		return
	}

	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Blocked", rows_updated)
}
