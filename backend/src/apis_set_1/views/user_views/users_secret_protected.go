package user_views

import (
	"context"
	"fmt"
	"learn_go/src/database"
	"learn_go/src/my_modules"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const api_secret = "1234"

func deleteUsercache(uuid string, ctx context.Context) {
	_users_keys, err := database.REDIS_DB_CONNECTION.Keys(ctx, "users___uuid="+uuid+"___/api/user/*").Result()
	if err == nil {
		for _, key := range _users_keys {
			database.REDIS_DB_CONNECTION.Del(ctx, key)
			log.WithFields(log.Fields{
				"key": key,
			}).Debugln(">>>>>>>>>>>>>>>> Redis, " + key + " Removed")
		}
	}
}

func InvalidateUsercache(c *gin.Context) {
	if c.GetHeader("secret") != api_secret {
		my_modules.CreateAndSendResponse(c, http.StatusForbidden, "error", "Invalid secret", nil)
		return
	}
	ctx := context.Background()
	database.REDIS_DB_CONNECTION.Set(ctx, "users_update_in_progress", "true", time.Second*0)

	// erasing pagination caches
	_paginated_keys, err := database.REDIS_DB_CONNECTION.Keys(ctx, "users___paginated*").Result()
	if err == nil {
		for _, key := range _paginated_keys {
			database.REDIS_DB_CONNECTION.Del(ctx, key)
			log.WithFields(log.Fields{
				"key": key,
			}).Debugln(">>>>>>>>>>>>>>>> Redis, users___paginated removed")
		}
	}

	db_connection := database.POSTGRES_DB_CONNECTION
	db_query := `SELECT uuid FROM users WHERE id=$1`
	rows, err := db_connection.Query(ctx, db_query, c.Param("id"))
	if err != nil {
		if err != context.Canceled {
			log.WithFields(log.Fields{
				"error": err,
				"query": db_query,
			}).Errorln("QueryRow failed ==>")
		}
		my_modules.CreateAndSendResponse(c, http.StatusBadRequest, "error", "No record found", nil)
		return
	} else {
		defer rows.Close() //importent to prevent connection leak
		for rows.Next() {
			var uuid string
			err := rows.Scan(&uuid)
			if err != nil {
				log.Errorln(fmt.Sprintf("Scan failed: %v\n", err))
				continue
			}
			// Erasing single user detail from cache
			deleteUsercache(uuid, ctx)
		}
		if err := rows.Err(); err != nil {
			if err != context.Canceled {
				log.Errorln(fmt.Sprintf("Row Err in rows.Next/rows.Scan failed: %v\n", err))
			}
			my_modules.CreateAndSendResponse(c, http.StatusInternalServerError, "error", "Error in retriving user data", nil)
			return
		}
	}

	database.REDIS_DB_CONNECTION.Del(ctx, "users_update_in_progress")
	log.Infoln("deleted users_update_in_progress")

	var count int
	rows2:=database.POSTGRES_DB_CONNECTION.QueryRow(ctx,"SELECT COUNT(*) FROM users")
	err2:=rows2.Scan(&count)
	if err2==nil{
		log.Infoln(fmt.Sprintf("count=%d",count))
		err:=database.REDIS_DB_CONNECTION.Set(ctx,"users_count",count,time.Second*0).Err()
		if err!=nil{
			log.WithFields(log.Fields{
				"errors":err,
			}).Errorln("Error in setting user count to redis")
		}
	} else {
		log.WithFields(log.Fields{
			"errors":err2,
		}).Errorln("Error in getting user count")
	}

}
