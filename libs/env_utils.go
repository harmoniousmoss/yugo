package libs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file and retrieves a specific key
func LoadEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Key %s not found in .env", key)
	}
	return value
}
