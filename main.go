package main

import (
	"log"

	"github.com/KevinDanae/twitterClone/bd"
	"github.com/KevinDanae/twitterClone/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if bd.CheckConnection() == 0 {
		log.Fatal("Error connecting to database")
		return
	}

	handlers.Handlers()

}
