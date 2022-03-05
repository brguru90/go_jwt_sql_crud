package my_modules

import (
	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	log.SetLevel(log.DebugLevel)
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})
	log.SetReportCaller(true)
}
