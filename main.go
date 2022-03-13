package main

import (
	"fmt"
	"github.com/mhthrh/Api-Dockerized/API"
	"net/http"
)

func main() {
	fmt.Println("run server")
	server := http.Server{
		Addr:              ":8585",
		Handler:           &API.RequestHandler{},
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
	}
	server.ListenAndServe()
}
