package user

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/user"
	"log"
	"strconv"
	"strings"
)

func (c *UserUserCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	var err error

	arguments := strings.SplitN(args, " ", 2)
	idx, err := strconv.ParseUint(arguments[0], 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	var userData user.User
	err = json.Unmarshal([]byte(arguments[1]), &userData)
	if err != nil {
		log.Println("wrong args", arguments[1])
		return
	}

	err = (*c.userService).Update(idx, userData)
	if err != nil {
		log.Printf("fail to delete product with idx %d: %v", idx, err)
		return
	}

	outputMsgText := fmt.Sprintf("User %v was updated", arguments[0])

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMsgText,
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("UserUserCommander.Delete: error sending reply message to chat - %v", err)
	}
}
