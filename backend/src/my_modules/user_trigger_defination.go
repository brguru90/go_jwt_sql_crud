package my_modules

import (
	"context"
	"fmt"
	"learn_go/src/database"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func FindUsersForcacheInvalidate(user_id string)  {
	db_connection := database.POSTGRES_DB_CONNECTION
	ctx:=context.Background()
	db_query := `SELECT uuid FROM users WHERE id=$1`
	rows, err := db_connection.Query(ctx, db_query,user_id)
	if err != nil {
		if err != context.Canceled {
			log.WithFields(log.Fields{
				"error": err,
				"query": db_query,
			}).Errorln("QueryRow failed ==>")
		}
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
			go InvalidateCache(uuid)
		}
		if err := rows.Err(); err != nil {
			if err != context.Canceled {
				log.Errorln(fmt.Sprintf("Row Err in rows.Next/rows.Scan failed: %v\n", err))
			}
			return
		}
	}
}

func deleteUserCache(uuid string, ctx context.Context) {
	// Deletes the cache for the specified user by his ID
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

func eraseAllUserPaginationCache(ctx context.Context) {
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
}

func getUsersCount(ctx context.Context) {
	var count int
	rows:=database.POSTGRES_DB_CONNECTION.QueryRow(ctx,"SELECT COUNT(*) FROM users")
	err2:=rows.Scan(&count)
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

func modifyCacheProgressStatus(operation string, ctx context.Context) {
	const max_users_update_in_progress_ttl=time.Minute*5
	
	users_update_in_progress, err := database.REDIS_DB_CONNECTION.Get(ctx, "users_update_in_progress").Result()
	if err == nil {
		users_update_in_progress_int, _ := strconv.ParseInt(users_update_in_progress, 10, 64)
		if operation=="delete"{
			users_update_in_progress_int--
			if users_update_in_progress_int==0{
				database.REDIS_DB_CONNECTION.Del(ctx, "users_update_in_progress")
				log.Debugln("deleted users_update_in_progress")				
			}
		} else{
			users_update_in_progress_int++
		}

		// log.WithFields(log.Fields{
		// 	"users_update_in_progress_int":users_update_in_progress_int,
		// }).Debugln("modifyCacheProgressStatus")

		if users_update_in_progress_int!=0{
			database.REDIS_DB_CONNECTION.Set(ctx, "users_update_in_progress", strconv.FormatInt(users_update_in_progress_int, 10),max_users_update_in_progress_ttl )
		}
	} else {
		if operation != "delete" {
			database.REDIS_DB_CONNECTION.Set(ctx, "users_update_in_progress", "1", max_users_update_in_progress_ttl)
		}
	}
}

func InvalidateCache(uuid string) {
	log.WithFields(log.Fields{
		"uuid":uuid,
	}).Debugln("invalidateCache....")
	ctx := context.Background()

	modifyCacheProgressStatus("insert",ctx)
	eraseAllUserPaginationCache(ctx)
	deleteUserCache(uuid, ctx)
	modifyCacheProgressStatus("delete",ctx)

	getUsersCount(ctx)
}


