package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiSpecificMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("apiSpecificMiddleware ===>", c.Request.URL.Path)
		// Before calling handler
		c.Next()
		// After calling handler
	}
}

func hello_api(c *gin.Context) {
	bodyAsByteArray, _ := ioutil.ReadAll(c.Request.Body)
	jsonBody := string(bodyAsByteArray)
	fmt.Println("Request body: \n", jsonBody)

	// work for 1st level depth
	// m := map[string]string{}
	var m interface{}

	err := json.Unmarshal([]byte(bodyAsByteArray), &m)
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Printf("type=%T, value=%v\n", m, m)
		// fmt.Println("test val==>", m["name"])
	}

	c.String(http.StatusOK, "Welcome")
}
