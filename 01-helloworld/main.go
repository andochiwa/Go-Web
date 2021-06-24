package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(response http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintln(response, "hello world", request.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)

	// 创建路由
	_ = http.ListenAndServe(":8080", nil)
}
