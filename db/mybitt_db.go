package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbUri string = "root:123456@tcp(127.0.0.1:3306)/mybitt_data"

func CreateMyBittDbConnection() *sql.DB {
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
