package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "请求地址:", r.URL.Path)
	fmt.Fprintln(w, "请求地址后查询的字符串:", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的信息:", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息:", r.Header["Accept-Encoding"])
}

func main() {
	http.HandleFunc("/", handler)

	_ = http.ListenAndServe("localhost:8080", nil)
}
