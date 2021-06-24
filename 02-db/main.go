package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goweb/02-db/model"
)

func main() {
	user := model.User{Username: "mary", Email: "aaa@g.com"}
	err := user.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}
