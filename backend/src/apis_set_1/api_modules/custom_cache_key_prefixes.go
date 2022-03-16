package api_modules

import (
	"learn_go/src/my_modules"
	"time"

	"github.com/gin-gonic/gin"
)

func ForUserPagination(c *gin.Context) string {
	if c.Query("page") != "" {
		// admin uuid could have been return,if access level is implemented
		return "paginated"
	}
	payload, ok := my_modules.ExtractTokenPayload(c)
	if !ok {
		// log.Panicln("Error in extracting payload from token")
		return "uuid=" + time.Now().String()
	}
	return "uuid=" + payload.Data.UUID
}
