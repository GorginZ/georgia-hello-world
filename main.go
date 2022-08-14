package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	metadata "github.com/GorginZ/georgia-hello-world/metadata"
	"github.com/GorginZ/georgia-hello-world/routes"
)

func main() {
	router()
}

func router() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", routes.HandleRoot).Methods("GET")
	r.HandleFunc("/status", routes.HandleStatus).Methods("GET")

	listen_port, found := os.LookupEnv("HTTP_PORT")
	if !found {
		listen_port = "8001"
		log.Printf("HTTP_PORT not found in env using default: %s", listen_port)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", listen_port),
		Handler: r,
	}
	log.Printf("Version: %s", metadata.Version)
	log.Printf("Sha: %s", metadata.Sha)
	log.Printf("Description: %s", metadata.Description)
	log.Printf("Listening on: %s", listen_port)

	log.Fatal(srv.ListenAndServe())
}
