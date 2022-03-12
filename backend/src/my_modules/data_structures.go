package my_modules

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
)

// structure member first letter should be capital indorser to export it as json
// binding encforce required condition
// empty interface data type allow unknown json format for which we don't have structure defined
// empty interface converted to map object
type UserRow struct {
	Column_id          int64       `json:"id" binding:"required"`
	Column_uuid        string      `json:"uuid" binding:"required"`
	Column_email       string      `json:"email" binding:"required"`
	Column_name        string      `json:"name" binding:"required"`
	Column_description string      `json:"description" binding:"required"`
	Column_createdAt   interface{} `json:"createdAt"`
	Column_updatedAt   interface{} `json:"updatedAt"`
}

type NewUserRow struct {
	Column_id          int64       `json:"id"`
	Column_uuid        string      `json:"uuid"`
	Column_email       string      `json:"email" binding:"required"`
	Column_name        string      `json:"name" binding:"required"`
	Column_description string      `json:"description" binding:"required"`
	Column_createdAt   interface{} `json:"createdAt"`
	Column_updatedAt   interface{} `json:"updatedAt"`
}

type ActiveSessionsRow struct {
	Column_id        int64       `json:"id"`
	Column_uuid      string      `json:"uuid"`
	Column_user_uuid string      `json:"user_uuid"`
	Column_token_id  string      `json:"token_id" binding:"required"`
	Column_ua        string      `json:"ua"`
	Column_ip        string      `json:"ip"`
	Column_exp       int64       `json:"exp" binding:"required"`
	Column_status    string      `json:"status"`
	Column_createdAt interface{} `json:"createdAt"`
	Column_updatedAt interface{} `json:"updatedAt"`
}

type ResponseCacheStruct struct {
	ResponseData   string    `json:"response_data" binding:"required"`
	ContentType    string    `json:"content_type" binding:"required"`
	HTTPStatusCode int       `json:"http_status_code" binding:"required"`
	LastModified   time.Time `json:"last_modified" binding:"required"`
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}
