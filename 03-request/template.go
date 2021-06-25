package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handlerTemplate(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("03-request/template.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = files.Execute(w, "hello world")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", handlerTemplate)

	_ = http.ListenAndServe("localhost:8080", nil)
}
