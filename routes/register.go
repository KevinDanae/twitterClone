package routes

import (
	"encoding/json"
	"net/http"

	"github.com/KevinDanae/twitterClone/bd"
	"github.com/KevinDanae/twitterClone/models"
)

// Register is a function that registers a new user in to the database
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error registering user: "+err.Error(), 400) // 400 Bad Request, if the JSON is not correct
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Error registering user: Email is required", 400) // 400 Bad Request, if the JSON is not correct
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Error registering user: Password min 6 characters", 400) // 400 Bad Request, if the JSON is not correct
		return
	}

	_, exist, _ := bd.CheckUserExist(user.Email)
	if exist {
		http.Error(w, "Error registering user: User already exists", 400) // 400 Bad Request, if the JSON is not correct
		return
	}

	_, status, err := bd.RegisterUser(user)
	if err != nil {
		http.Error(w, "Error registering user: "+err.Error(), 400) // 400 Bad Request, if the JSON is not correct
		return
	}
	if !status {
		http.Error(w, "Error registering user: not register user", 400) // 400 Bad Request, if the JSON is not correct
		return
	}

	w.WriteHeader(http.StatusCreated)
}
