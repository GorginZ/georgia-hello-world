package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	router()
}

func router() {
	http.HandleFunc("/", handleRoot)

	listen_port, found := os.LookupEnv("HTTP_PORT")
	if !found {
		listen_port = "8001"
	}
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", listen_port),
	}
	log.Fatal(srv.ListenAndServe())
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
