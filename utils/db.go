package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

func init() {
	DB, err = sql.Open("mysql", "root:gjk0722@tcp(127.0.0.1:3306)/bookstore")
	if err != nil {
		panic(err.Error())
	}
}
