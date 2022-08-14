package routes

import (
	"encoding/json"
	"net/http"

	"github.com/GorginZ/georgia-hello-world/metadata"
)

type AppData struct {
	Metadata Metadata `json:"my-application"`
}

type Metadata struct {
	Version     string `json:"version"`
	Description string `json:"description"`
	Sha         string `json:"sha"`
}

func HandleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	meta := Metadata{Version: metadata.Version, Description: metadata.Description, Sha: metadata.Sha}
	appData := AppData{Metadata: meta}
	json.NewEncoder(w).Encode(appData)
}
