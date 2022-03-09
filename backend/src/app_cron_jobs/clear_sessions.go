package app_cron_jobs

import (
	"context"
	"fmt"
	"learn_go/src/database"
	"time"

	log "github.com/sirupsen/logrus"
)

func ClearExpiredToken() {
	_time := time.Now()
	log.WithFields(log.Fields{
		"time": _time.UnixMilli(),
	}).Debug(" -- ClearExpiredToken Cron job started -- ")

	db_connection := database.POSTGRES_DB_CONNECTION
	const sql_stmt string = `DELETE FROM active_sessions WHERE exp<=$1`
	res, err := db_connection.Exec(context.Background(), sql_stmt, _time.UnixMilli())
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"sql": fmt.Sprintf(`DELETE FROM active_sessions WHERE exp<=%v`, _time.UnixMilli()),
		}).Errorln("Failed to delete user data")
		return
	}

	log.WithFields(log.Fields{
		"Total_cleared_session_entry": res.RowsAffected(),
	}).Debug(" -- ClearExpiredToken Cron finished -- ")
}
