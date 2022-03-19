package my_modules

import (
	"learn_go/src/configs"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {

	if configs.EnvConfigs.APP_ENV == "production" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			PadLevelText:           true,
			ForceColors:            true,
		})
	}

	log.SetReportCaller(true)
}
