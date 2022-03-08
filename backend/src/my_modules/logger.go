package my_modules

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	log.SetLevel(log.DebugLevel)

	if os.Getenv("APP_ENV") == "production" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			PadLevelText:           true,
			ForceColors:            true,
		})
	}

	log.SetReportCaller(true)
}
