package main

import (
	"log"
	"net/http"

	"github.com/OMSV1/common"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	httpAddr := common.EnvString("HTTP_ADDR", ":3000")

	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("Starting server on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
