package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	httpAddr := "0.0.0.0:9999"
	if port := os.Getenv("PORT"); port != "" {
		httpAddr = "0.0.0.0:" + port
	}
	fmt.Println("HTTP server listening on ", httpAddr)

	log.Fatal(http.ListenAndServe(httpAddr, setupRouter()))
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", HealthCheckHandler).Methods("GET")
	return r
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HIT: healthcheck")
	_, _ = io.WriteString(w, "ALIVE")
}
