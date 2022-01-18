package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/KevinDanae/twitterClone/middlewares"
	"github.com/KevinDanae/twitterClone/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckBd(routes.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckBd(routes.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
