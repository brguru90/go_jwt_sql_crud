package database

import (
	"context"
	"fmt"
	"learn_go/src/configs"
	"os"
	"strings"

	"github.com/jackc/pgtype"
	pgtypeuuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	log "github.com/sirupsen/logrus"
)

// Postgres Connection is concurrent safe, so no need to lock while using
var POSTGRES_DB_CONNECTION *pgxpool.Pool

func ConnectPostgres() *pgxpool.Pool {
	// https://github.com/jackc/pgx

	var DB_USER string = configs.EnvConfigs.POSTGRES_DB_USER
	var DB_PASSWORD string = configs.EnvConfigs.POSTGRES_DB_PASSWORD
	var DB_HOST string = configs.EnvConfigs.POSTGRES_DB_HOST
	var DATABASE string = configs.EnvConfigs.POSTGRES_DATABASE
	var DB_PORT int64 = configs.EnvConfigs.POSTGRES_DB_PORT

	// trying to connect database assuming specified database is already present
	var DB_URL string = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DATABASE)
	dbconfig, err := pgxpool.ParseConfig(DB_URL)

	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"DB_URL": DB_URL,
		}).Error("Unable to connect to database ==>")
	}

	dbconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		conn.ConnInfo().RegisterDataType(pgtype.DataType{
			Value: &pgtypeuuid.UUID{},
			Name:  "uuid",
			OID:   pgtype.UUIDOID,
		})
		return nil
	}

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
			// Since specified DB not present, connecting postgres DB to create a new Database
			var POSTGRES_URL string = fmt.Sprintf("postgresql://%s:%s@%s:%d/postgres", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT)
			log.Infoln(fmt.Sprintf("Connecting to %s", POSTGRES_URL))
			_db_connection, err := pgx.Connect(context.Background(), POSTGRES_URL)
			if err != nil {
				log.WithFields(log.Fields{
					"error":  err,
					"DB_URL": DB_URL,
				}).Fatalln("Unable to connect to Database ==>")
				os.Exit(1)
			}

			// Creating specified DB
			log.Infoln(fmt.Sprintf("Creating %s Database ", DATABASE))
			var create_db string = fmt.Sprintf("CREATE DATABASE %s;", DATABASE)
			_rows, err2 := _db_connection.Query(context.Background(), create_db)
			if err2 != nil {
				log.WithFields(log.Fields{
					"error": err2,
					"query": create_db,
				}).Fatalln("Unable to create Database ==>")
				os.Exit(1)
			}
			defer _rows.Close()
			_db_connection.Close(context.Background())
		}

		{
			// reconnecting to specified DB, after the DB is created
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

	// Creating the Table schemas
	InitUserModels(dbpool)
	InitActiveSessionsModels(dbpool)

	POSTGRES_DB_CONNECTION = dbpool

	return dbpool
}
