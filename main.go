package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jimmyvallejo/blog-aggregator-go/internal/api/v1/handlers"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	APICfg := APIConfig{
		Port: port,
	}

	mux.HandleFunc("GET /v1/healthz", handlers.HandlerReadiness)
	mux.HandleFunc("GET /v1/err", handlers.HandlerError)

	srv := &http.Server{
		Addr:    ":" + APICfg.Port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", APICfg.Port)
	log.Fatal(srv.ListenAndServe())

}
