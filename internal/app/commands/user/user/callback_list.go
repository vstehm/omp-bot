package user

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Cursor uint64 `json:"cursor"`
	Limit  uint64 `json:"limit"`
}

func (c *UserUserCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("UserUserCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	outputMsgText := fmt.Sprintf("Here list of users starting from %v, limit %v: \n\n", parsedData.Cursor, parsedData.Limit)

	extendedLimit := parsedData.Limit + 1 //to get next page start
	users, err := (*c.userService).List(parsedData.Cursor, extendedLimit)
	if err != nil {
		outputMsgText = err.Error()
	}
	var serializedData []byte

	if uint64(len(users)) == extendedLimit {
		lastId := users[len(users)-1].Id
		users = users[:parsedData.Limit]
		serializedData, _ = json.Marshal(CallbackListData{
			Cursor: lastId,
			Limit:  parsedData.Limit,
		})
	}
	for _, u := range users {
		outputMsgText += u.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	if len(serializedData) > 0 {
		newCallbackPath := path.CallbackPath{
			Domain:       "user",
			Subdomain:    "user",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", newCallbackPath.String()),
			),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("UserUserCommander.List: error sending reply message to chat - %v", err)
	}
}
