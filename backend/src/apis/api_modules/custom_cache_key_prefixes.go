package api_modules

import (
	"learn_go/src/my_modules"

	"github.com/gin-gonic/gin"
)

func ForUserPagination(c *gin.Context) string {
	if c.Query("page") != "" {
		payload, ok := my_modules.ExtractTokenPayload(c)
		if !ok {
			return ""
		}
		return "uuid=" + payload.Data.UUID
	}
	return ""
}
