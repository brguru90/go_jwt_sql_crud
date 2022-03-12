package apis_set_1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func test_api(c *gin.Context) {
	c.String(http.StatusOK, "the param sent %s", c.Param("id"))
}
