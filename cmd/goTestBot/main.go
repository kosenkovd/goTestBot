package main

import (
	"log"
	"os"

	"github.com/kosenkovd/goTestBot/internal/app/commands"
	"github.com/kosenkovd/goTestBot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

const tokenVariable = "TG_TEST_TOKEN"

func main() {
	godotenv.Load()

	token := os.Getenv(tokenVariable)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
		Offset:  0,
		Limit:   0,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		commander.HandleUpdate(update)
	}
}
