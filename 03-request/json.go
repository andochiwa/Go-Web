package main

import (
	"encoding/json"
	"fmt"
	"goweb/02-db/model"
	"net/http"
)

func handlerJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := model.User{Id: 1, Username: "admin"}
	// 将user转换成json格式
	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = w.Write(userJson)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", handlerJson)

	_ = http.ListenAndServe("localhost:8080", nil)
}
