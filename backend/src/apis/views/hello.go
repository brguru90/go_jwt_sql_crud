package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Hello_api(c *gin.Context) {
	bodyAsByteArray, _ := ioutil.ReadAll(c.Request.Body)
	jsonBody := string(bodyAsByteArray)
	log.Infoln("Request body: \n", jsonBody)

	// work for 1st level depth
	// m := map[string]string{}
	var m interface{}

	err := json.Unmarshal([]byte(bodyAsByteArray), &m)
	if err != nil {
		log.Errorln("err", err)
	} else {
		log.Infoln(fmt.Sprintf("type=%T, value=%v\n", m, m))
		// fmt.Println("test val==>", m["name"])
	}

	c.String(http.StatusOK, "Welcome hello")
}
