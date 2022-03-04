package my_modules

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateResponsePayload(status string, message string, Data interface{}) string {
	var responseBody ResponseBody
	responseBody.Data = Data
	responseBody.Status = status
	responseBody.Message = message

	result, _ := json.MarshalIndent(responseBody, "", "  ")
	return string(result)
}

func CreateAndSendResponse(c *gin.Context, HTTP_Status int, status string, message string, Data interface{}) {

	if HTTP_Status == 0 {
		HTTP_Status = http.StatusOK
	}
	if Data == nil {
		c.JSON(HTTP_Status, gin.H{"status": status, "message": message})
		return
	}

	var responseBody ResponseBody
	responseBody.Data = Data
	responseBody.Status = status
	responseBody.Message = message

	result, _ := json.MarshalIndent(responseBody, "", "  ")

	c.Writer.Header().Set("Content-Type", "application/json")
	c.String(HTTP_Status, string(result))
}
