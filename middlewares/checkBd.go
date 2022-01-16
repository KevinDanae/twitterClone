package middlewares

import (
	"net/http"

	"github.com/KevinDanae/twitterClone/bd"
)

// CheckBd is a middleware that checks database connection
func CheckBd(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Cannot connect to database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
