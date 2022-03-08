package views

import (
	"context"
	"encoding/json"
	"learn_go/src/database"
	"learn_go/src/my_modules"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func SignUp(c *gin.Context) {
	db_connection := database.POSTGRES_DB_CONNECTION
	// in Progress

	var newUserRow NewUserRow
	if err := c.ShouldBindJSON(&newUserRow); err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Invalid input data format", nil)
		return
	}
	_time := time.Now()
	const sql_stmt string = `INSERT INTO users ("name","email","description","uuid","createdAt","updatedAt") VALUES($1,$2,$3,$4,$5,$6) RETURNING id,uuid,name,email,description,"createdAt","updatedAt"`
	err := db_connection.QueryRow(context.Background(), sql_stmt, newUserRow.Column_name, newUserRow.Column_email, newUserRow.Column_description, uuid.New().String(), _time, _time).Scan(&newUserRow.Column_id, &newUserRow.Column_uuid, &newUserRow.Column_name, &newUserRow.Column_email, &newUserRow.Column_description, &newUserRow.Column_createdAt, &newUserRow.Column_updatedAt)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Error("Error in inserting user data")
		my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in Regestering new user", nil)
		return
	}

	token_payload, _ := json.Marshal(map[string]string{
		"email": newUserRow.Column_email,
		"uuid":  newUserRow.Column_uuid,
	})
	access_token, access_token_payload := my_modules.GenerateAccessToken(
		newUserRow.Column_email,
		my_modules.EnsureCsrfToken(c),
		string(token_payload),
	)

	newUserRow_json, _ := json.Marshal(newUserRow)
	my_modules.SetCookie(c, "access_token", access_token, access_token_payload.Exp, true)
	my_modules.SetCookie(c, "user_data", string(newUserRow_json), access_token_payload.Exp, true)

	my_modules.CreateAndSendResponse(c, http.StatusOK, "success", "Regesteration successfully", newUserRow)
}

func Login(c *gin.Context) {

	c.String(http.StatusOK, "Welcome hello")
}

func LoginStatus(c *gin.Context) {
	decoded_token, err, http_status := my_modules.LoginStatus(c)
	if err != "" {
		my_modules.CreateAndSendResponse(c, http_status, "error", err, nil)
		return
	}
	if decoded_token != nil {
		my_modules.CreateAndSendResponse(c, http_status, "success", "active", nil)
	}
}
