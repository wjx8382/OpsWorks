package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 设置路由处理函数
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/api/get", handleGet)
	http.HandleFunc("/api/post", handlePost)

	// 启动HTTP服务器
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start HTTP server: ", err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET request handled")
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST request handled")
}
