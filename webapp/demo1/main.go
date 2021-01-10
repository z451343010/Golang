package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello World", r.URL.Path)

}

func main() {
	// HandleFunc 帮助转换函数
	// 可以将一个带有正确签名的函数 f
	// 转换成带有方法 f 的 Handler
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
