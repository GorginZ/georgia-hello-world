package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var Version = "development"

type AppData struct {
	Metadata Metadata `json:"my-application"`
}

type Metadata struct {
	Version     string `json:"version"`
	Description string `json:"description"`
	Sha         string `json:"sha"`
}

func main() {
	router()
}

func getVersion() string {
	return Version
}
func getSha() string {
	return "abc53458585"
}
func getDesc() string {
	return "text"
}

func router() {
	// currently is returning root for all other paths
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/status", handleStatus)

	listen_port, found := os.LookupEnv("HTTP_PORT")
	if !found {
		listen_port = "8001"
	}
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", listen_port),
	}
	log.Printf("Version: %s", Version)

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

func handleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	version := getVersion()
	description := getDesc()
	sha := getSha()

	meta := Metadata{Version: version, Description: description, Sha: sha}
	appData := AppData{Metadata: meta}
	json.NewEncoder(w).Encode(appData)
	w.WriteHeader(200)
}
