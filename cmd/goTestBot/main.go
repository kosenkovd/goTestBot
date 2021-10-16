package main

import (
	"log"
	"os"

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

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "help":
			processHelp(bot, update.Message)
		default:
			processDefault(bot, update.Message)
		}
	}
}

func processHelp(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Help: Some help.")
	bot.Send(msg)
}

func processDefault(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: \""+inputMessage.Text+"\"")
	msg.ReplyToMessageID = inputMessage.MessageID

	bot.Send(msg)
}
