package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func test_api(c *gin.Context) {
	c.String(http.StatusOK, "the param sent %s", c.Param("id"))
}
