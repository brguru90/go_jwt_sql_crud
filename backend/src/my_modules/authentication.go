package my_modules

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"learn_go/src/database"
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

type TokenPayload struct {
	Email string `json:"email" binding:"required"`
	UUID  string `json:"uuid" binding:"required"`
}
type AccessToken struct {
	Data TokenPayload `json:"data" binding:"required"`
	// Data         interface{}  `json:"data" binding:"required"`
	Uname        string `json:"uname" binding:"required"`
	Token_id     string `json:"token_id" binding:"required"`
	Exp          int64  `json:"exp" binding:"required"`
	IssuedAtTime int64  `json:"issued_at" binding:"required"`
	Csrf_token   string `json:"csrf_token" binding:"required"`
}

type AccessTokenClaims struct {
	jwt.StandardClaims
	AccessToken `json:"accessToken" binding:"required"`
}

func randomBytes(size int) (blk []byte, err error) {
	blk = make([]byte, size)
	_, err = rand.Read(blk)
	return
}

func GenerateAccessToken(uname string, csrf_token string, data TokenPayload) (string, AccessTokenClaims) {
	time_now := time.Now().UnixMilli()
	token_id := ""

	if _rand, r_err := randomBytes(100); r_err == nil {
		token_id = data.UUID + "_" + base64.StdEncoding.EncodeToString(_rand) + "_" + string(time_now)
	}

	var accessTokenPayload AccessTokenClaims = AccessTokenClaims{
		AccessToken: AccessToken{
			Uname:        uname,
			Token_id:     token_id,
			Data:         data,
			IssuedAtTime: time_now,
			Exp:          time_now + JWT_TOKEN_EXPIRE,
		},
	}

	// accessTokenPayload.Uname = uname
	// accessTokenPayload.Token_id = token_id
	// accessTokenPayload.Data = data
	// accessTokenPayload.IssuedAtTime = time_now
	// accessTokenPayload.Exp = time_now + JWT_TOKEN_EXPIRE

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenPayload)
	token_string, token_err := token.SignedString([]byte(JWT_SECRET_KEY))

	if token_err != nil {
		log.WithFields(log.Fields{
			"token_err":      token_err,
			"token_data":     accessTokenPayload,
			"JWT_SECRET_KEY": JWT_SECRET_KEY,
		}).Panic("Error in signing JWT")
	}

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

func DeleteCookie(c *gin.Context, key string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     key,
		MaxAge:   -1,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
	})
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

func Authenticate(c *gin.Context, newUserRow NewUserRow) AccessToken {
	token_payload := TokenPayload{
		Email: newUserRow.Column_email,
		UUID:  newUserRow.Column_uuid,
	}
	access_token, access_token_payload := GenerateAccessToken(
		newUserRow.Column_email,
		EnsureCsrfToken(c),
		token_payload,
	)
	newUserRow_json, _ := json.Marshal(newUserRow)
	SetCookie(c, "access_token", access_token, access_token_payload.Exp, true)
	SetCookie(c, "user_data", string(newUserRow_json), access_token_payload.Exp, true)
	return access_token_payload.AccessToken
}

func LoginStatus(c *gin.Context) (AccessToken, string, int, bool) {
	var token_claims AccessTokenClaims
	access_token, err := c.Cookie("access_token")

	if err != nil {
		if err == http.ErrNoCookie {
			return AccessToken{}, "No access_token Cookie present", http.StatusBadRequest, false
		}
		log.WithFields(log.Fields{
			"Error": err,
		}).Error("Unknown error in extracting the cookie")
		return AccessToken{}, "Unknown error in extracting the cookie", http.StatusInternalServerError, false
	}

	token, err := jwt.ParseWithClaims(access_token, &token_claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		e, ok := err.(*jwt.ValidationError)
		if err == jwt.ErrSignatureInvalid {
			return AccessToken{}, "Invalid token signature", http.StatusForbidden, false
		}
		log.WithFields(log.Fields{
			"Error":        err,
			"access_token": access_token,
			"e":            e,
			"ok":           ok,
			"e.Errors":     e.Errors,
		}).Error("Unknown error in Decrypting token")
		return AccessToken{}, "Unknown error in Decrypting token", http.StatusInternalServerError, false
	}
	if !token.Valid {
		return AccessToken{}, "unAuthorized", http.StatusForbidden, false
	}

	decoded_token_json, _ := json.Marshal(token_claims.AccessToken)
	_, r_err := database.REDIS_DB_CONNECTION.Get(context.Background(), string(decoded_token_json)).Result()
	if r_err == nil {
		return token_claims.AccessToken, "Session blocked", http.StatusForbidden, false
	}
	return token_claims.AccessToken, "", http.StatusOK, true
}

func ExtractTokenPayload(c *gin.Context) (TokenPayload, bool) {
	c_data, ok := c.Get("decoded_token")
	if !ok {
		CreateAndSendResponse(c, http.StatusOK, "error", "Not able to get user data from current session", nil)
		c.Abort()
		return TokenPayload{}, false
	}
	return c_data.(AccessToken).Data, true
}
