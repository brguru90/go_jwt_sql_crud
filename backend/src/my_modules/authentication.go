package my_modules

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AccessToken struct {
	Data       interface{} `json:"data" binding:"required"`
	Uname      string      `json:"uname" binding:"required"`
	Token_id   string      `json:"token_id" binding:"required"`
	Exp        int64       `json:"exp" binding:"required"`
	Iat        int64       `json:"iat" binding:"required"`
	Csrf_token string      `json:"csrf_token" binding:"required"`
}

var JWT_TOKEN_EXPIRE,_=strconv.ParseInt(os.Getenv("JWT_TOKEN_EXPIRE"), 10, 64) 
var JWT_SECRET_KEY string=os.Getenv("JWT_SECRET_KEY")

func GenerateAccessToken(uname string, csrf_token string, data interface{}) (string,jwt.MapClaims) {
	time_now:=time.Now().UnixMilli()
	access_token_payload:=jwt.MapClaims{}
	access_token_payload["uname"] = uname
	access_token_payload["data"] = data
	access_token_payload["csrf_token"] = csrf_token
	access_token_payload["iat"] = time_now
	access_token_payload["exp"]=time_now+JWT_TOKEN_EXPIRE

	token:=jwt.NewWithClaims(&jwt.SigningMethodHMAC{},access_token_payload)
	token_string,token_err:=token.SignedString(JWT_SECRET_KEY)

	if token_err!=nil{
		log.WithFields(log.Fields{
			"token_err":token_err,
			"token_data":access_token_payload,
		}).Panic("Error in signing JWT")
	}

	return token_string,access_token_payload

}

func setCookie(c *gin.Context,key string,value string,expires int64,httpOnly bool)  {
	
	c.Set()
}
