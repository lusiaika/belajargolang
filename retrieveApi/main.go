package main

import (
	"fmt"
	"net/http"
	"retrieveApi/handler"
	"retrieveApi/middleware"
	"time"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/users-url", handler.DataUser)

	var handler http.Handler = mux
	handler = middleware.MiddlewareAuth(handler)
	handler = middleware.MiddlewareAllowOnlyGet(handler)

	server := new(http.Server)
	server.Addr = "127.0.0.1:8000"
	server.Handler = handler
	server.WriteTimeout = 15 * time.Second
	server.ReadTimeout = 15 * time.Second

	fmt.Println("server started at localhost:8000")
	server.ListenAndServe()
}
