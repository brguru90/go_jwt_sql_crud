package app_cron_jobs

import (
	"context"
	"fmt"
	"learn_go/src/database"
	"time"

	log "github.com/sirupsen/logrus"
)

func ClearExpiredToken() {
	// pgxpool is a concurrency-safe connection pool for pgx. ignore --race errors

	ctx := context.Background()
	_time := time.Now()

	log.WithFields(log.Fields{
		"time": _time,
	}).Debug(" -- ClearExpiredToken Cron job started -- ")

	const sql_stmt string = `DELETE FROM active_sessions WHERE exp<=$1`
	res, err := database.POSTGRES_DB_CONNECTION.Exec(ctx, sql_stmt, _time.UnixMilli())
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
