package user

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/user"
	"log"
)

func (c *UserUserCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	var err error
	var newUsers []uint64
	var userDataSlice []user.User

	err = json.Unmarshal([]byte(args), &userDataSlice)
	if err != nil {
		log.Println("wrong args for slice", args)

		var userData user.User
		err = json.Unmarshal([]byte(args), &userData)
		if err != nil {
			log.Println("wrong args for user", args)
			return
		}
		userDataSlice = append(userDataSlice, userData)
	}

	for _, u := range userDataSlice {
		id, err := c.userService.Create(u)
		if err != nil {
			log.Printf("fail to delete product with idx %d: %v", id, err)
		}
		newUsers = append(newUsers, id)
	}

	if len(newUsers) == 0 {
		log.Println("No users was created")
		return
	}
	outputMsgText := fmt.Sprintf("User %v was created", newUsers)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMsgText,
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("UserUserCommander.Delete: error sending reply message to chat - %v", err)
	}
}
