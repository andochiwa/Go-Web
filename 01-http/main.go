package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
//func handler(response http.ResponseWriter, request *http.Request) {
//	_, _ = fmt.Fprintln(response, "hello world", request.URL.Path)
//}

// 使用自定义struct创建处理器方法
type MyHandler struct {
}

func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "custom handler")
}

func main() {
	myHandler := MyHandler{}
	http.Handle("/", &myHandler)

	// 创建路由
	_ = http.ListenAndServe(":8080", nil)
}
