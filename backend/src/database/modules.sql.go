package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

func Exec_SQL(db_connection *pgxpool.Pool, SQL string) {
	_, err := db_connection.Exec(context.Background(), SQL)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"SQL":   SQL,
		}).Errorln("Query failed ==>")
		return
	} else {
		log.WithFields(log.Fields{
			"SQL": SQL,
		}).Infoln("SQL successfully executed ==>")
	}
}
