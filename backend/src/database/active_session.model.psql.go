package database

import (
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

const create_active_sessions_table string = `CREATE TABLE IF NOT EXISTS "active_sessions" ("id"   SERIAL , "uuid" UUID, "user_uuid" VARCHAR(255), "token_id" VARCHAR(255) NOT NULL UNIQUE, "ua" TEXT, "ip" VARCHAR(255), "exp" BIGINT, "status" VARCHAR(255), "createdAt" TIMESTAMP WITH TIME ZONE NOT NULL, "updatedAt" TIMESTAMP WITH TIME ZONE NOT NULL, PRIMARY KEY ("id"));`

const active_sessions_uuid_token_id string = `CREATE UNIQUE INDEX "active_sessions_uuid_token_id" ON "active_sessions" ("uuid", "token_id");`

const active_sessions_user_uuid string = `CREATE INDEX "active_sessions_user_uuid" ON "active_sessions" ("user_uuid");`

const active_sessions_token_id string = `CREATE INDEX "active_sessions_token_id" ON "active_sessions" ("token_id");`

func InitActiveSessionsModels(db_connection *pgxpool.Pool) {
	log.Info("InitActiveSessionsModels")

	Exec_SQL(db_connection, create_active_sessions_table)
	Exec_SQL(db_connection, active_sessions_uuid_token_id)
	Exec_SQL(db_connection, active_sessions_user_uuid)
	Exec_SQL(db_connection, active_sessions_token_id)
}
