package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KevinDanae/twitterClone/bd"
	"github.com/KevinDanae/twitterClone/jwt"
	"github.com/KevinDanae/twitterClone/models"
)

// Login is a function that logs a user in to the database, endpoints: /login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "User or password invalid"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	document, exist := bd.Login(t.Email, t.Password)
	if !exist {
		http.Error(w, "User or password invalid", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "An error ocurred"+err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
