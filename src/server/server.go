package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server : Struct to hold server relevant data
type Server struct {
	Router *mux.Router
	Logger *os.File
	Port   string
}

// CreateServer : Defaultconstructor
func CreateServer() *Server {
	return &Server{
		Router: CreateRouter(),
		Logger: CreateLogger(),
		Port:   ":8080",
	}
}

// Start : Start the server
func (s *Server) Start() {
	server := http.Server{
		Addr:    s.Port,
		Handler: handlers.LoggingHandler(s.Logger, s.Router),
	}

	fmt.Println("Running")
	server.ListenAndServe()
}
