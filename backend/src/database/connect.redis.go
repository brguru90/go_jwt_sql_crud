package database

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

// Redis Connection is concurrent safe, so no need to lock while using
var REDIS_DB_CONNECTION *redis.Client

func ConnectRedis() {
	// https://github.com/go-redis/redis
	// https://github.com/go-redis/redis/issues/166

	log.Info("Connecting to Redis....")
	var ctx = context.Background()
	redis_conn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_ping := redis_conn.Ping(ctx)
	// checking that is it possible to write data to database
	err := redis_conn.Set(ctx, "test_connection", "value", 5*time.Minute).Err()
	if err != nil {
		log.WithFields(log.Fields{
			"REDIS_DB_CONNECTION": redis_conn,
			"_ping":               _ping,
			"Error":               err,
		}).Panic("Unable to connect redis")
		return
	}

	log.WithFields(log.Fields{
		"REDIS_DB_CONNECTION": redis_conn,
		"Ping":                _ping,
	}).Info("Redis database connected successfully")

	REDIS_DB_CONNECTION = redis_conn
}
