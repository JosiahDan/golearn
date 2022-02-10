package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method:", r.Method)
	// /go
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// 回复
	var result []string = ["123", "48586"]
w.Write([]byte(result))
}

func main() {
	http.HandleFunc("/go", myHandler)
	http.ListenAndServe("127.0.0.1:9876", nil)
}
