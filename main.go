package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/massimomarsiglia/cs-skins-market-models/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()
}


