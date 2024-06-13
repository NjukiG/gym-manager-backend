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

	err := godotenv.Load(filepath.Join(pwd, ".env"))

	if err != nil {
		log.Fatal("FAiled to load .env file...")
	}
}
