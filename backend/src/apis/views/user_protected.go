package views

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Row struct {
	COL_id          int64       `json:"id"`
	COL_uuid        string      `json:"uuid"`
	COL_email       string      `json:"email"`
	COL_name        string      `json:"name"`
	COL_description string      `json:"description"`
	COL_createdAt   interface{} `json:"createdAt"`
	COL_updatedAt   interface{} `json:"updatedAt"`
}

func GetUserData(c *gin.Context, db_connection *pgxpool.Pool) {

	var uuid string = c.Query("uuid")
	if uuid != "" {
		var db_query string = fmt.Sprintf(`select * from users where uuid='%s'; `, uuid)
		rows, err := db_connection.Query(context.Background(), db_query)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\tdb_query==>%s\n", err, db_query)
			c.JSON(http.StatusOK, gin.H{"status": "error", "message": "No record found"})
			return
		} else {
			defer rows.Close()
			var rowSlice []Row
			for rows.Next() {
				var r Row
				err := rows.Scan(&r.COL_id, &r.COL_uuid, &r.COL_name, &r.COL_email, &r.COL_description, &r.COL_createdAt, &r.COL_updatedAt)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
				}
				rowSlice = append(rowSlice, r)
			}
			fmt.Printf("type=%T\nresult=%v", rowSlice, rowSlice)

			if err := rows.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "Row Err failed: %v\n", err)
			}

			result, _ := json.MarshalIndent(rowSlice, "", "  ")
			c.Writer.Header().Set("Content-Type", "application/json")
			c.String(200, fmt.Sprintf(`{"data":%s, "msg": "Found","status": "success"}`, string(result)))
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Didn't got UUID"})
		return
	}
}

func UpdateUserData(c *gin.Context, db_connection *pgxpool.Pool) {

	c.String(http.StatusOK, "Welcome hello")
}

func GetActiveSession(c *gin.Context, db_connection *pgxpool.Pool) {

	c.String(http.StatusOK, "Welcome hello")
}

func Deleteuser(c *gin.Context, db_connection *pgxpool.Pool) {

	c.String(http.StatusOK, "Welcome hello")
}

func Logout(c *gin.Context, db_connection *pgxpool.Pool) {

	c.String(http.StatusOK, "Welcome hello")
}

func BlockSession(c *gin.Context, db_connection *pgxpool.Pool) {

	c.String(http.StatusOK, "Welcome hello")
}
