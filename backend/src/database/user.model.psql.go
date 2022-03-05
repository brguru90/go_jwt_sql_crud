package database

import (
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

const create_user_table string = `CREATE TABLE IF NOT EXISTS "users" ("id"   SERIAL , "uuid" UUID, "name" VARCHAR(255) NOT NULL, "email" VARCHAR(255) NOT NULL UNIQUE, "description" TEXT, "createdAt" TIMESTAMP WITH TIME ZONE NOT NULL, "updatedAt" TIMESTAMP WITH TIME ZONE NOT NULL, PRIMARY KEY ("id"));`

const users_email_uuid string = `CREATE UNIQUE INDEX "users_email_uuid" ON "users" ("email", "uuid");`

const users_uuid string = ` CREATE INDEX "users_uuid" ON "users" ("uuid");`

func InitUserModels(db_connection *pgxpool.Pool) {
	log.Info("InitUserModels")

	Exec_SQL(db_connection, create_user_table)
	Exec_SQL(db_connection, users_email_uuid)
	Exec_SQL(db_connection, users_uuid)
}
