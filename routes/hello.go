package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}
	resp := make(map[string]string)
	resp["message"] = "Hello World!"
	json, err := json.Marshal(resp)
	if err != nil {
		log.Printf("JSON marshal Error: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(json)
}
