package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
}

// Load loads the environment variables from the .env file.
func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		BotToken: os.Getenv("BOT_TOKEN"),
	}
}
