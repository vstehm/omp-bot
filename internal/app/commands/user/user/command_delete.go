package user

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *UserUserCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	var err error

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	result, err := (*c.userService).Remove(idx)
	if err != nil {
		log.Printf("fail to delete product with idx %d: %v", idx, err)
		return
	}

	outputMsgText := fmt.Sprintf("%v", result)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMsgText,
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("UserUserCommander.Delete: error sending reply message to chat - %v", err)
	}
}
