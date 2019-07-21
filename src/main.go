package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	err := http.ListenAndServe(":8083", mux)
	if err != nil {
		fmt.Println("启动失败")
	}
}
