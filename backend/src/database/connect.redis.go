package database

import (
	"context"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var REDIS_DB_CONNECTION *redis.Client

func ConnectRedis() {
	log.Info("Connecting to Redis....")

	var ctx = context.Background()
	REDIS_DB_CONNECTION = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_ping := REDIS_DB_CONNECTION.Ping(ctx)
	err := REDIS_DB_CONNECTION.Set(ctx, "test_connection", "value", 0).Err()
	if err != nil {
		log.WithFields(log.Fields{
			"REDIS_DB_CONNECTION": REDIS_DB_CONNECTION,
			"_ping":               _ping,
			"Error":               err,
		}).Panic("Unable to connect redis")
		return
	}

	log.WithFields(log.Fields{
		"REDIS_DB_CONNECTION": REDIS_DB_CONNECTION,
		"Ping":                _ping,
	}).Info("Redis database connected successfully")
}
