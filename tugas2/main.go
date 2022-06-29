package main

import (
	"net/http"
	"time"
	"tugas2/handler"

	"github.com/gorilla/mux"
)

func main() {

	s := mux.NewRouter()
	h := handler.NewUserHandler()
	s.HandleFunc("/user/{id}", h.UsersHandler)
	s.HandleFunc("/user", h.UsersHandler)
	srv := &http.Server{
		Handler: s,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()

}
