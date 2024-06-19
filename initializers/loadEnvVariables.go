package initializers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {

	pwd, err1 := os.Getwd()

	if err1 != nil {
		panic(err1)
	}

	log.Printf("Current working directory: %s", pwd)

	err := godotenv.Load(filepath.Join(pwd, ".env"))

	if err != nil {
		log.Fatal("Failed to load .env file...")
	}
}
