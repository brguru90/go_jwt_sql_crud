package database

import (
	"fmt"
	"path/filepath"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

const create_user_table string = `CREATE TABLE IF NOT EXISTS "users" ("id"   SERIAL , "uuid" UUID, "name" VARCHAR(255) NOT NULL, "email" VARCHAR(255) NOT NULL UNIQUE, "description" TEXT, "createdAt" TIMESTAMP WITH TIME ZONE NOT NULL, "updatedAt" TIMESTAMP WITH TIME ZONE NOT NULL, PRIMARY KEY ("id"));`

const users_email_uuid string = `CREATE UNIQUE INDEX "users_email_uuid" ON "users" ("email", "uuid");`

const users_uuid string = ` CREATE INDEX "users_uuid" ON "users" ("uuid");`

func get_trigger_sqls() []string {
	// * chmod o+rx $HOME, if permission denied
	user_update_trigger, err := filepath.Abs("src/database/triggers/user_update_trigger.so")
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Panicln("not able to get absolute path of user_update_trigger.so")
	}

	// Creating the function & mapping the function to the executable build
	var users_triggers_map_func = fmt.Sprintf(`CREATE FUNCTION user_update_trigger()
RETURNS TRIGGER AS '%s'
LANGUAGE C;`, user_update_trigger)

	// Creating the trigger & mapping function to user table
	var users_triggers_register_to_table = `CREATE TRIGGER user_update_trigger_on_insert
AFTER INSERT OR UPDATE
ON users
FOR EACH ROW
EXECUTE PROCEDURE user_update_trigger();`

	// ? To view the trigger events
	// SELECT * FROM information_schema.triggers;

	return []string{users_triggers_map_func, users_triggers_register_to_table}
}

func InitUserModels(db_connection *pgxpool.Pool) {

	log.Info("InitUserModels")

	Exec_SQL(db_connection, create_user_table)
	Exec_SQL(db_connection, users_email_uuid)
	Exec_SQL(db_connection, users_uuid)

	for _, sql := range get_trigger_sqls() {
		Exec_SQL(db_connection, sql)
	}
}
