package views

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}


// @BasePath /api
// @Summary Test API
// @Schemes
// @Description just to test,and work space to play with request
// @Tags Hello
// @Accept json
// @Produce plain
// @Param page path int false "page"
// @Param limit path int false "limit"
// @Success 200 {object} my_modules.ResponseFormat
// @Failure 500 {object} my_modules.ResponseFormat
// @Router /hello/{page}/{limit} [get]
func Hello_api(c *gin.Context) {
	// jsonBody := string(bodyAsByteArray)
	// log.Infoln("Request body: \n", jsonBody)

	// // work for 1st level depth
	// // m := map[string]string{}
	// var m interface{}

	// err := json.Unmarshal([]byte(bodyAsByteArray), &m)
	// if err != nil {
	// 	log.Errorln("err", err)
	// } else {
	// 	log.Infoln(fmt.Sprintf("type=%T, value=%v\n", m, m))
	// 	// fmt.Println("test val==>", m["name"])
	// }

	_raw_dt, _ := c.GetRawData()
	_params, _ := json.Marshal(c.Params)

	fmt.Println("c.GetRawData()", string(_raw_dt))
	fmt.Println("c.Request.RequestURI", c.Request.RequestURI)
	fmt.Println("c.Request.URL", c.Request.URL)
	fmt.Println("c.Params", string(_params))
	fmt.Println("c.FullPath()", c.FullPath())
	fmt.Println("c.Request.URL.RawPath", c.Request.URL.RawPath)
	fmt.Println("c.Request.URL.RawQuery", c.Request.URL.RawQuery)
	fmt.Println("c.Request.URL.RawFragment", c.Request.URL.RawFragment)
	fmt.Println("c.Request.PostForm", c.Request.PostForm)

	h := sha1.New()
	h.Write([]byte(c.Request.RequestURI + string(_params) + string(_raw_dt)))

	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	fmt.Println("hash=",string(h.Sum(nil)))

	c.String(http.StatusOK, "Welcome hello")

	// fmt.Println("Response body: " + blw.body.String())
	// fmt.Println("c.Request.Response.Body Content-Type", c.Writer.Header().Get("Content-Type"))
	// fmt.Println("c.Request.Response.StatusCode", c.Writer.Status())
}
