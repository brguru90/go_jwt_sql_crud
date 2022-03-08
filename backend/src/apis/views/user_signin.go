package views

import (
	"context"
	"learn_go/src/database"
	"learn_go/src/my_modules"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

func SignUp(c *gin.Context) {
	db_connection := database.POSTGRES_DB_CONNECTION
	// in Progress

	var newUserRow my_modules.NewUserRow
	if err := c.ShouldBindJSON(&newUserRow); err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Invalid input data format", nil)
		return
	}
	_time := time.Now()
	const user_sql_stmt string = `INSERT INTO users ("name","email","description","uuid","createdAt","updatedAt") VALUES($1,$2,$3,$4,$5,$6) RETURNING id,uuid,name,email,description,"createdAt","updatedAt"`
	user_err := db_connection.QueryRow(context.Background(), user_sql_stmt, newUserRow.Column_name, newUserRow.Column_email, newUserRow.Column_description, uuid.New().String(), _time, _time).Scan(&newUserRow.Column_id, &newUserRow.Column_uuid, &newUserRow.Column_name, &newUserRow.Column_email, &newUserRow.Column_description, &newUserRow.Column_createdAt, &newUserRow.Column_updatedAt)
	if user_err != nil {
		log.WithFields(log.Fields{
			"Error": user_err,
		}).Error("Error in inserting user data")
		my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in Regestering new user", nil)
		return
	}

	access_token_payload := my_modules.Authenticate(c, newUserRow)

	var active_sessions_id int64
	var active_sessions_uuid string
	_time = time.Now()

	const active_sessions_sql_stmt string = `INSERT INTO active_sessions ("uuid","user_uuid","token_id","ua","ip","exp","status","createdAt","updatedAt") VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id,uuid`
	active_sessions_err := db_connection.QueryRow(context.Background(), active_sessions_sql_stmt, uuid.New().String(), newUserRow.Column_uuid, access_token_payload.Token_id, c.GetHeader("User-Agent"), c.ClientIP(), access_token_payload.Exp, "active", _time, _time).Scan(&active_sessions_id, &active_sessions_uuid)
	if active_sessions_err != nil {
		log.WithFields(log.Fields{
			"Error": active_sessions_err,
		}).Error("Error in inserting active_sessions data")
		my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in Regestering new user while marking active", nil)
		return
	}

	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Regesteration successfully", newUserRow)
}

type UserEmailID struct {
	Email string `json:"email" binding:"required"`
}

func Login(c *gin.Context) {
	db_connection := database.POSTGRES_DB_CONNECTION

	var userEmailID UserEmailID
	if err := c.ShouldBindJSON(&userEmailID); err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Invalid input data format", nil)
		return
	}

	var newUserRow my_modules.NewUserRow
	const sql_stmt string = `SELECT * FROM users WHERE email=$1`
	err := db_connection.QueryRow(context.Background(), sql_stmt, userEmailID.Email).Scan(&newUserRow.Column_id, &newUserRow.Column_uuid, &newUserRow.Column_name, &newUserRow.Column_email, &newUserRow.Column_description, &newUserRow.Column_createdAt, &newUserRow.Column_updatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.WithFields(log.Fields{
				"Error": err,
				"Email": userEmailID.Email,
			}).Warning("Error in finding user email")
			my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Invalid credential", nil)
			return
		}
		log.WithFields(log.Fields{
			"Error": err,
			"Email": userEmailID.Email,
		}).Error("Error in finding user email")
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Error in finding user", nil)
		return
	}

	access_token_payload := my_modules.Authenticate(c, newUserRow)

	var active_sessions_id int64
	var active_sessions_uuid string
	_time := time.Now()

	const active_sessions_sql_stmt string = `INSERT INTO active_sessions ("uuid","user_uuid","token_id","ua","ip","exp","status","createdAt","updatedAt") VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id,uuid`
	active_sessions_err := db_connection.QueryRow(context.Background(), active_sessions_sql_stmt, uuid.New().String(), newUserRow.Column_uuid, access_token_payload.Token_id, c.GetHeader("User-Agent"), c.ClientIP(), access_token_payload.Exp, "active", _time, _time).Scan(&active_sessions_id, &active_sessions_uuid)
	if active_sessions_err != nil {
		log.WithFields(log.Fields{
			"Error": active_sessions_err,
		}).Error("Error in inserting active_sessions data")
		my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in Loggin in user while marking active", nil)
		return
	}

	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Authorization success", newUserRow)
}

func LoginStatus(c *gin.Context) {
	_, err, http_status, ok := my_modules.LoginStatus(c)
	if err != "" {
		my_modules.CreateAndSendResponse(c, http_status, "error", err, nil)
		return
	}
	if ok {
		my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "active", nil)
	}
}
