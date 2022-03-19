package configs

import (
	"encoding/json"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type ENV_CONFIGS struct {
	SERVER_PORT                int64
	POSTGRES_DB_USER           string
	POSTGRES_DB_PASSWORD       string
	POSTGRES_DB_HOST           string
	POSTGRES_DATABASE          string
	POSTGRES_DB_PORT           int64
	JWT_SECRET_KEY             string
	JWT_TOKEN_EXPIRE_IN_MINS   int64
	ENABLE_REDIS_CACHE         bool
	RESPONSE_CACHE_TTL_IN_SECS int64
	APP_ENV                    string
	NODE_ENV                   string
	GIN_MODE                   string
	DISABLE_COLOR              bool
}

var EnvConfigs ENV_CONFIGS

func strToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func InitEnv() {
	EnvConfigs = ENV_CONFIGS{
		SERVER_PORT:                strToInt64(os.Getenv("SERVER_PORT")),
		POSTGRES_DB_USER:           os.Getenv("POSTGRES_DB_USER"),
		POSTGRES_DB_PASSWORD:       os.Getenv("POSTGRES_DB_PASSWORD"),
		POSTGRES_DB_HOST:           os.Getenv("POSTGRES_DB_HOST"),
		POSTGRES_DATABASE:          os.Getenv("POSTGRES_DATABASE"),
		POSTGRES_DB_PORT:           strToInt64(os.Getenv("POSTGRES_DB_PORT")),
		JWT_SECRET_KEY:             os.Getenv("JWT_SECRET_KEY"),
		JWT_TOKEN_EXPIRE_IN_MINS:   strToInt64(os.Getenv("JWT_TOKEN_EXPIRE_IN_MINS")),
		ENABLE_REDIS_CACHE:         os.Getenv("ENABLE_REDIS_CACHE") == "true",
		RESPONSE_CACHE_TTL_IN_SECS: strToInt64(os.Getenv("RESPONSE_CACHE_TTL_IN_SECS")),
		APP_ENV:                    os.Getenv("APP_ENV"),
		NODE_ENV:                   os.Getenv("NODE_ENV"),
		GIN_MODE:                   os.Getenv("GIN_MODE"),
		DISABLE_COLOR:              os.Getenv("DISABLE_COLOR") == "true",
	}

	a, _ := json.MarshalIndent(EnvConfigs, "\t", "")
	log.Infoln("ENV_CONFIGS=" + string(a))
}
