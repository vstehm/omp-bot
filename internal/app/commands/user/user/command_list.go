package user

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
	"strconv"
	"strings"
)

const (
	Cursor uint64 = 1
	Limit  uint64 = 20
)

func (c *UserUserCommander) List(inputMessage *tgbotapi.Message) {
	var cursor, limit = Cursor, Limit
	var err error

	argsS := inputMessage.CommandArguments()

	args := strings.Fields(argsS)

	if len(args) == 2 {
		cursor, err = strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Println("wrong args", args[0])
			return
		}
		limit, err = strconv.ParseUint(args[1], 10, 64)
		if err != nil {
			log.Println("wrong args", args[1])
			return
		}
	}

	outputMsgText := fmt.Sprintf("Here list of users starting from %v, limit %v: \n\n", cursor, limit)

	extendedLimit := limit + 1 //to get next page start

	users, err := c.userService.List(cursor, extendedLimit)
	if err != nil {
		outputMsgText = err.Error()
	}

	var serializedData []byte
	if uint64(len(users)) == extendedLimit {
		//remove cursor user
		lastId := users[len(users)-1].Id
		users = users[:limit]
		serializedData, _ = json.Marshal(CallbackListData{
			Cursor: lastId,
			Limit:  limit,
		})
	}

	for _, u := range users {
		outputMsgText += u.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	if len(serializedData) > 0 {
		callbackPath := path.CallbackPath{
			Domain:       "user",
			Subdomain:    "user",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("UserUserCommander.List: error sending reply message to chat - %v", err)
	}
}
