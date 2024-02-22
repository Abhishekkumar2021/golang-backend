package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// This file contains all the constants used in the application
const DB_NAME = "go_api"

func LoadEnv()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
