package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
)

func init() {
	database, err := sql.Open("mysql", "root:root@/go_test")
	if err != nil {
		panic(err.Error())
	}
	Db = database
}
