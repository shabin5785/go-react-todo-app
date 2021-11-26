package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "todo"
)

var db *sql.DB
var err error

func SetupPostgres() {

	host := "localhost"

	h := os.Getenv("pghost")

	if len(h) > 0 {
		host = h
	}

	fmt.Printf("\n Connecting to %s\n", host)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("DB connected")

}

func GetAllTodoItems(c *gin.Context) {
	rows, err := db.Query("select * from list;")
	b := checkUnknownError(c, err)
	if b {
		return
	}

	items := make([]Item, 0)
	if rows != nil && !b {
		defer rows.Close()
		for rows.Next() {
			item := Item{}
			err = rows.Scan(&item.ID, &item.Item, &item.Done)
			item.Item = strings.TrimSpace(item.Item)
			b = checkUnknownError(c, err)
			items = append(items, item)
		}
	}
	if !b {
		c.JSON(http.StatusOK, gin.H{"items": items})
	}

}

func UpdateTodoItem(c *gin.Context) {

	var item Item

	err = c.BindJSON(&item)
	b := checkUnknownError(c, err)
	if !b {
		b = checkUpdateFields(item)
		if !b {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid update fields"})
		} else {
			var exists bool
			err = db.QueryRow("select * from list where id=$1;", item.ID).Scan(&exists)
			if !checkNoRowError(c, err) {
				fmt.Println("> ", item.Done)
				_, err = db.Query("update list set item=$1 , done=$2 where id=$3;", item.Item, item.Done, item.ID)
				if err != nil {
					fmt.Println(err)
					c.JSON(http.StatusInternalServerError, gin.H{"message": "update failed"})
				} else {
					c.JSON(http.StatusOK, gin.H{"item": item})
				}

			}
		}
	}

}

func CreateTodoItem(c *gin.Context) {

	var item Item

	err = c.BindJSON(&item)
	b := checkUnknownError(c, err)
	if !b {
		if len(strings.TrimSpace(item.Item)) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid item"})
			return
		}

		row, err := db.Query("insert into list(item, done) values($1,$2) RETURNING id;", item.Item, false)
		b = checkUnknownError(c, err)

		if !b {
			row.Next()
			row.Scan(&item.ID)
			c.JSON(http.StatusCreated, gin.H{"item": item})
		}
	}
}

func DeleteTodoItem(c *gin.Context) {
	id := c.Param("id")

	if len(strings.TrimSpace(id)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id"})
	} else {
		var exists bool
		err = db.QueryRow("select * from list where id=$1;", id).Scan(&exists)
		if !checkNoRowError(c, err) {
			_, err = db.Query("delete from list where id=$1;", id)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "delete failed"})
			} else {
				c.JSON(http.StatusOK, gin.H{"id": id})
			}
		}
	}
}

func checkUnknownError(c *gin.Context, err error) bool {
	b := false
	if err != nil {
		b = true
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
	}
	return b
}

func checkNoRowError(c *gin.Context, err error) bool {
	b := false
	if err != nil && err == sql.ErrNoRows {
		fmt.Println(err)
		b = true
		c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
	}
	return b
}

func checkUpdateFields(item Item) bool {
	if len(strings.TrimSpace(item.Item)) < 1 {
		return false
	} else if len(strings.TrimSpace(item.OldItem)) < 1 {
		return false
	} else {
		return true
	}
}
