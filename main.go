package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const defaultPort = ":3000"

var notImplemented = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Not Implemented"))
	})

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("port")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()

	router.HandleFunc("/status", notImplemented).Methods("GET")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	recoveryRouter := handlers.RecoveryHandler()(loggedRouter)


	log.Fatal(http.ListenAndServe(":3000", recoveryRouter))
}
