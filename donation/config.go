package donation

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Webhhok string
	Secret  string
}

func GetEnv() (config Config) {
	if err := godotenv.Load(); err != nil {
		log.Println("Can't load .env file on the root directory.")
	}

	webhook := os.Getenv("OMISE_WEBHOOK")
	secret := os.Getenv("OMISE_SECRET")

	return Config{Webhhok: webhook, Secret: secret}
}
