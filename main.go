package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	router()
}

func router() {
	// currently is returning root for all other paths
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
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}
	w.WriteHeader(200)
	resp := make(map[string]string)
	resp["message"] = "Hello World!"
	json, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("JSON marshal Error. Err: %s", err)
	}
	w.Write(json)
}
