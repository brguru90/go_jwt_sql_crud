package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

func ConnectPostgres() *pgxpool.Pool {

	// https://github.com/jackc/pgx

	var DB_USER string = os.Getenv("DB_USER")
	var DB_PASSWORD string = os.Getenv("DB_PASSWORD")
	var DB_HOST string = os.Getenv("DB_HOST")
	var DATABASE string = os.Getenv("DATABASE")
	var DB_PORT string = os.Getenv("DB_PORT")

	var DB_URL string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DATABASE)

	dbpool, err := pgxpool.Connect(context.Background(), DB_URL)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"DB_URL": DB_URL,
		}).Fatalln("Unable to connect to database ==>")
		os.Exit(1)
	}
	return dbpool
}
