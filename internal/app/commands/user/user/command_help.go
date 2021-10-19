package user

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *UserUserCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__user__user — show help\n"+
			"/get__user__user — get a entity\n"+
			"/list__user__user — get a list of your entity\n"+
			"/delete__user__user — delete an existing entity\n"+
			"/new__user__user — create a new entity\n"+
			"/edit__user__user — edit a entity",
	)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("UserUserCommander.Help: error sending reply message to chat - %v", err)
	}
}
