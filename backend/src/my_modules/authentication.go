package my_modules

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var JWT_TOKEN_EXPIRE_IN_MINS, _ = strconv.ParseInt(os.Getenv("JWT_TOKEN_EXPIRE_IN_MINS"), 10, 64)
var JWT_TOKEN_EXPIRE = JWT_TOKEN_EXPIRE_IN_MINS * 60 * 1000
var JWT_SECRET_KEY string = os.Getenv("JWT_SECRET_KEY")

type AccessToken struct {
	Data       string `json:"data" binding:"required"`
	Uname      string `json:"uname" binding:"required"`
	Token_id   string `json:"token_id" binding:"required"`
	Exp        int64  `json:"exp" binding:"required"`
	IssuedAt   int64  `json:"issued_at" binding:"required"`
	Csrf_token string `json:"csrf_token" binding:"required"`
}

func randomBytes(size int) (blk []byte, err error) {
	blk = make([]byte, size)
	_, err = rand.Read(blk)
	return
}

func GenerateAccessToken(uname string, csrf_token string, data string) (string, AccessToken) {
	time_now := time.Now().UnixMilli()
	token_id := ""

	access_token_payload := jwt.MapClaims{}
	access_token_payload["uname"] = uname
	access_token_payload["data"] = data
	access_token_payload["token_id"] = token_id
	access_token_payload["csrf_token"] = csrf_token
	access_token_payload["issued_at"] = time_now
	access_token_payload["exp"] = time_now + JWT_TOKEN_EXPIRE

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, access_token_payload)
	token_string, token_err := token.SignedString([]byte(JWT_SECRET_KEY))

	if token_err != nil {
		log.WithFields(log.Fields{
			"token_err":      token_err,
			"token_data":     access_token_payload,
			"JWT_SECRET_KEY": JWT_SECRET_KEY,
		}).Panic("Error in signing JWT")
	}

	var accessTokenPayload AccessToken
	accessTokenPayload.Uname = uname
	accessTokenPayload.Token_id = token_id
	accessTokenPayload.Data = data
	accessTokenPayload.IssuedAt = time_now
	accessTokenPayload.Exp = time_now + JWT_TOKEN_EXPIRE

	return token_string, accessTokenPayload

}

func SetCookie(c *gin.Context, key string, value string, expires int64, httpOnly bool) {
	// c.SetSameSite(http.SameSiteStrictMode)
	// c.SetCookie(key, value, int(expires), "/", "", false, httpOnly)
	_time := time.UnixMilli(expires)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(value),
		Expires:  _time,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
		HttpOnly: httpOnly,
	})
}

func LoginStatus(c *gin.Context) (interface{}, string, int) {
	decoded_token := jwt.MapClaims{}

	access_token, err := c.Cookie("access_token")

	if err != nil {
		if err == http.ErrNoCookie {
			return nil, "No access_token Cookie present", http.StatusBadRequest
		}
		log.WithFields(log.Fields{
			"Error": err,
		}).Error("Unknown error in extracting the cookie")
		return nil, "Unknown error in extracting the cookie", http.StatusInternalServerError
	}

	token, err := jwt.ParseWithClaims(access_token, &decoded_token, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		e, ok := err.(*jwt.ValidationError)
		if err == jwt.ErrSignatureInvalid {
			return nil, "Invalid token signature", http.StatusForbidden
		}
		log.WithFields(log.Fields{
			"Error":        err,
			"access_token": access_token,
			"e":            e,
			"ok":           ok,
			"e.Errors":     e.Errors,
		}).Error("Unknown error in Decrypting token")
		return nil, "Unknown error in Decrypting token", http.StatusInternalServerError
	}
	if !token.Valid {
		return nil, "Invalid token", http.StatusForbidden
	}

	return decoded_token, "", http.StatusOK
}

func ValidateCredential(c *gin.Context) {

}

func EnsureCsrfToken(c *gin.Context) string {
	var csrf_token string = ""
	if _rand, r_err := randomBytes(100); r_err == nil {
		csrf_token = base64.StdEncoding.EncodeToString(_rand)
		c.Header("csrf_token", csrf_token)
	} else {
		log.WithFields(log.Fields{
			"err": r_err,
		}).Error("Error in generating csrf token")
	}
	return csrf_token
}
