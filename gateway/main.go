package main

import (
	"log"
	"net/http"
)

const (
	httpAddr = ":9090"
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	if err := http.ListenAndServe(httpAddr, mux); err != nil{
		log.Fatal("Failed to start http server")
	}

}