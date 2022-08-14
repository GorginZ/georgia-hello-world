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
	w.WriteHeader(200)
	resp := make(map[string]string)
	resp["message"] = "Hello World!"
	json, err := json.Marshal(resp)
	if err != nil {
		//I don't want it to die
		log.Fatalf("JSON marshal Error. Err: %s", err)
	}
	w.Write(json)
}
