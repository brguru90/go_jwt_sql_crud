package database

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

var POSTGRES_DB_CONNECTION *pgxpool.Pool

func ConnectPostgres() {

	// https://github.com/jackc/pgx

	var DB_USER string = os.Getenv("DB_USER")
	var DB_PASSWORD string = os.Getenv("DB_PASSWORD")
	var DB_HOST string = os.Getenv("DB_HOST")
	var DATABASE string = os.Getenv("DATABASE")
	var DB_PORT string = os.Getenv("DB_PORT")

	var DB_URL string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DATABASE)

	log.Infoln(fmt.Sprintf("Connecting to %s", DB_URL))
	dbpool, err := pgxpool.Connect(context.Background(), DB_URL)
	if err != nil {
		if !strings.Contains(fmt.Sprintf("%v", err), fmt.Sprintf(`database "%s" does not exist`, DATABASE)) {
			log.WithFields(log.Fields{
				"error":  err,
				"DB_URL": DB_URL,
			}).Fatalln("Unable to connect to database ==>")
			os.Exit(1)
		}
		log.WithFields(log.Fields{
			"error":  err,
			"DB_URL": DB_URL,
		}).Warningln("Database doesn't exists")

		{
			var POSTGRES_URL string = fmt.Sprintf("postgresql://%s:%s@%s:%s/postgres", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT)
			log.Infoln(fmt.Sprintf("Connecting to %s", POSTGRES_URL))
			_db_connection, err := pgx.Connect(context.Background(), POSTGRES_URL)
			if err != nil {
				log.WithFields(log.Fields{
					"error":  err,
					"DB_URL": DB_URL,
				}).Fatalln("Unable to connect to Database ==>")
				os.Exit(1)
			}

			log.Infoln(fmt.Sprintf("Creating %s Database ", DATABASE))
			var create_db string = fmt.Sprintf("CREATE DATABASE %s;", DATABASE)
			_, err2 := _db_connection.Query(context.Background(), create_db)
			if err2 != nil {
				log.WithFields(log.Fields{
					"error": err2,
					"query": create_db,
				}).Fatalln("Unable to create Database ==>")
				os.Exit(1)
			}
			_db_connection.Close(context.Background())
		}

		{
			log.Infoln(fmt.Sprintf("Reconnecting to %s", DB_URL))
			dbpool, err = pgxpool.Connect(context.Background(), DB_URL)
			if err != nil {
				log.WithFields(log.Fields{
					"error":  err,
					"DB_URL": DB_URL,
				}).Fatalln("Unable to connect to database ==>")
				os.Exit(1)
			}
		}
	}

	InitUserModels(dbpool)
	InitActiveSessionsModels(dbpool)

	POSTGRES_DB_CONNECTION = dbpool
}
