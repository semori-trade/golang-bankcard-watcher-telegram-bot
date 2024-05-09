package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/semori-trade/golang-bankcard-watcher-telegram-bot/telegramBot"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// init is invoked before main()
func initBot(telegramToken string) {
	telegramBot.NewTelegramBot(telegramToken, telegramBot.Options{Debug: true})
}

func main() {
	telegramToken, exists := os.LookupEnv("TELEGRAM_TOKEN")

	if !exists {
		log.Fatal("TELEGRAM_TOKEN environment variable not set")
	}

	initBot(telegramToken)
}
