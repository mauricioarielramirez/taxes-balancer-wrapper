package main

import (
	"log"
	"net/http"

	"taxes-balancer-wrapper/pkg/infraestructure/delivery/server"
)

func main() {
	startServer()
}

func startServer() {
	s := server.New()

	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
