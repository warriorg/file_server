package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"server"
)

func main() { //设置访问的路由
	r := mux.NewRouter()
	r.HandleFunc("/", server.Hello)
	r.HandleFunc("/upload", server.Upload)
	r.HandleFunc("/{key}", server.Load)

	err := http.ListenAndServe(":9900", r) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		os.Exit(1)
	}
}
