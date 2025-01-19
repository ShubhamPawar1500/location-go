package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}
