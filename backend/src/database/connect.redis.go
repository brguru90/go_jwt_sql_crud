package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	redispool "github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

// Redis Connection is concurrent safe, so no need to lock while using
var REDIS_DB_CONNECTION *redis.Client
var REDIS_DB_CONNECTION_POOL *redispool.Pool

func ConnectRedis() {
	// !warning, this repo is not clear how it will handle connection pool, & when i ran benchmark i got very poor performance
	// ! So this library here its not used for implementing APIs caching
	// https://github.com/go-redis/redis
	// https://github.com/go-redis/redis/issues/166

	log.Info("Connecting to Redis....")
	var ctx = context.Background()
	REDIS_DB_CONNECTION = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_ping := REDIS_DB_CONNECTION.Ping(ctx)
	// checking that is it possible to write data to database
	err := REDIS_DB_CONNECTION.Set(ctx, "test_connection", "value", 5*time.Minute).Err()
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

type Person struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
}

func InitRedisPool() {
	// !Important, Using this library i got very good performance while running benchmarks,
	// ! So this library will mainly used for API caching
	// https://github.com/gomodule/redigo/
	REDIS_DB_CONNECTION_POOL = &redispool.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redispool.Conn, error) {
			conn, err := redispool.Dial("tcp", "localhost:6379")
			if err != nil {
				log.WithFields(log.Fields{
					"Error": err,
				}).Panic("Unable to connect redis pool")
			}
			return conn, err
		},
	}

	conn := REDIS_DB_CONNECTION_POOL.Get()
	defer conn.Close()
	_, err := redispool.String(conn.Do("PING"))
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Panic("Unable to ping redis ppool")
	}

	if err := RedisPoolSet("test_connection_pool", "value", 2*time.Minute); err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Panic("Unable to write into redis pool")
	}

	if _, err := RedisPoolGet("test_connection_pool"); err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Errorln("Unable to write into redis pool")
	}

	var person Person = Person{
		Name: "guru",
		Age:  27,
	}
	if err := RedisPoolSetJSON("person", person, 2*time.Minute); err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Panic("Unable to write into redis pool")
	}

	if err := RedisPoolGetJSON("person", &Person{}); err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Errorln("Unable to write into redis pool")
	}

}

func RedisPoolSet(key string, value string, ttl_sec time.Duration) error {
	conn := REDIS_DB_CONNECTION_POOL.Get()
	defer conn.Close()

	var err error
	if ttl_sec <= 0 {
		_, err = conn.Do("SET", key, value)
	} else {
		_, err = conn.Do("SET", key, value, "EX", fmt.Sprintf("%v", (ttl_sec).Seconds()))
	}
	return err
}

func RedisPoolGet(key string) (string, error) {
	conn := REDIS_DB_CONNECTION_POOL.Get()
	defer conn.Close()

	s, err := redispool.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return s, err
}

func RedisPoolSet_Bytes(key string, value []byte, ttl_sec time.Duration) error {
	conn := REDIS_DB_CONNECTION_POOL.Get()
	defer conn.Close()

	var err error
	if ttl_sec <= 0 {
		_, err = conn.Do("SET", key, value)
	} else {
		_, err = conn.Do("SET", key, value, "EX", fmt.Sprintf("%v", (ttl_sec).Seconds()))
	}
	return err
}

func RedisPoolGet_Bytes(key string) ([]byte, error) {
	conn := REDIS_DB_CONNECTION_POOL.Get()
	defer conn.Close()

	s, err := redispool.Bytes(conn.Do("GET", key))
	if err != nil {
		return []byte{}, err
	}
	return s, err
}

func RedisPoolSetJSON(key string, value interface{}, ttl_sec time.Duration) error {
	json_str, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return RedisPoolSet_Bytes(key, json_str, ttl_sec)
}

func RedisPoolGetJSON(key string, destination interface{}) error {
	val, err := RedisPoolGet_Bytes(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(val, &destination)
}
