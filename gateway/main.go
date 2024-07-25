package main

import (
	"log"
	"net/http"

	common "github.com/dunky-star/omsv2-common"
	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":9090")
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("Starting http server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil{
		log.Fatal("Failed to start http server")
	}

}