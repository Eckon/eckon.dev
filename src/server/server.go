package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
)

// Start : Starts the server
func Start() {
	router := CreateRouter()
	log := CreateLogger()

	server := http.Server{
		Addr:    ":8080",
		Handler: handlers.LoggingHandler(log, router),
	}

	fmt.Println("Running")
	server.ListenAndServe()
}
