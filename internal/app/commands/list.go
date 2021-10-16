package commands

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMessage := "All the products: \n\n"

	products := c.productService.List()

	for i, p := range products {
		outputMessage += fmt.Sprintf("%v. ", i+1)
		outputMessage += p.Title
		outputMessage += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessage)

	serializedData, _ := json.Marshal(
		CallbackData{
			Command: "list",
			Offset:  5,
			Length:  5,
		})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
		),
	)

	c.bot.Send(msg)
}
