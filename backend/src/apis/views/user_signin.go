package views

import (
	"encoding/json"
	"learn_go/src/my_modules"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	// in Progress

	var newUserRow NewUserRow
	if err := c.ShouldBindJSON(&newUserRow); err != nil {
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "Invalid input data format", nil)
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
	my_modules.SetCookie(c, "access_token", access_token, access_token_payload.Exp, true)

	c.String(http.StatusOK, "Signup")
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
