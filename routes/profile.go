package routes

import (
	"encoding/json"
	"net/http"

	"github.com/KevinDanae/twitterClone/bd"
)

// Profile is a function that consult the profile of a user
func Profile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Error searching profile: id is required", http.StatusBadRequest)
		return // 400 Bad Request, if the JSON is not correct
	}

	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Error searching profile: "+err.Error(), http.StatusBadRequest)
		return // 400 Bad Request, if the JSON is not correct
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
