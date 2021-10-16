package commands

import (
	"log"
	"strings"

	"github.com/kosenkovd/goTestBot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		callbackData := strings.Split(update.CallbackQuery.Data, "_")

		switch callbackData[0] {
		case "list":
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, callbackData[1])
			c.bot.Send(msg)
		}

		return
	}

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
}
