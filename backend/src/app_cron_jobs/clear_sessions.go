package app_cron_jobs

import (
	"context"
	"fmt"
	"learn_go/src/database"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

var clear_expire_token_mut sync.Mutex
var clear_expire_token_db_connection *pgxpool.Conn = nil

func ClearExpiredToken() {
	clear_expire_token_mut.Lock()
	defer clear_expire_token_mut.Unlock()
	ctx := context.Background()
	if clear_expire_token_db_connection == nil {
		var conn_err error
		clear_expire_token_db_connection, conn_err = database.POSTGRES_DB_CONNECTION.Acquire(ctx)
		if conn_err != nil {
			log.WithFields(log.Fields{
				"conn_err": conn_err,
				"sql":      "Error in acquiring connection from pool",
			}).Errorln("Failed to delete user data")
		}
	}

	_time := time.Now()
	log.WithFields(log.Fields{
		"time": _time,
	}).Debug(" -- ClearExpiredToken Cron job started -- ")

	const sql_stmt string = `DELETE FROM active_sessions WHERE exp<=$1`
	res, err := clear_expire_token_db_connection.Exec(ctx, sql_stmt, _time.UnixMilli())
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
	clear_expire_token_db_connection.Conn().Close(ctx)
	clear_expire_token_db_connection = nil
}
