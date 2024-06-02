package main

import (
	"log"
	"net/http"

	"calculator/internal/orchestrator"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	orchestrator.RegisterHandlers(r)
	http.Handle("/", r)

	log.Println("Starting orchestrator on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
