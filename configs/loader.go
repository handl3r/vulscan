package configs

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvironments() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Can not load env")
	}
}
