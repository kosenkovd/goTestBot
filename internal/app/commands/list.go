package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMessage := "All the products: \n\n"

	products := c.productService.List()

	for i, p := range products {
		outputMessage += fmt.Sprintf("%v. ", i+1)
		outputMessage += p.String()
		outputMessage += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessage)

	c.bot.Send(msg)
}
