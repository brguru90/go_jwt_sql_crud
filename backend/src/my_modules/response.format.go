package my_modules

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateResponsePayload(status string, message string, Data interface{}) string {

	responseBody := make(map[string]interface{})
	responseBody["data"] = Data
	responseBody["status"] = status
	responseBody["msg"] = message

	result, _ := json.MarshalIndent(responseBody, "", "  ")
	return string(result)
}

func canSendResponse(c *gin.Context) bool {
	if c_err := c.Request.Context().Err(); c_err != nil {
		if c_err == context.Canceled {
			log.Debugln(fmt.Sprintf(`Request cancelled for %v route`, c.FullPath()))
			c.Abort()
			return false
		}
		log.WithFields(log.Fields{
			"Path": c.FullPath(),
			"Err":  c_err,
		}).Error("Error before sending response")
		return false
	}
	return true
}

func CreateAndSendResponse(c *gin.Context, HTTP_Status int, status string, message string, Data interface{}) {

	// some time we may get runtime error when we trying to send response after API is cancelled
	// to avoid that, checking the context is cancelled
	if !canSendResponse(c) {
		c.JSON(http.StatusInternalServerError, gin.H{"status": status, "error": "aborted"})
		return
	}

	if HTTP_Status == 0 {
		HTTP_Status = http.StatusOK
	}
	if Data == nil {
		if canSendResponse(c) {
			c.JSON(HTTP_Status, gin.H{"status": status, "msg": message})
		}
		return
	}

	// responseBody := make(map[string]interface{})
	// responseBody["data"] = Data
	// responseBody["status"] = status
	// responseBody["msg"] = message

	responseBody:=ResponseFormat{
		Status: status,
		Msg: message,
		Data: Data,
	}

	result, _ := json.MarshalIndent(responseBody, "", "  ")

	if canSendResponse(c) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(HTTP_Status, string(result))
	}
}
