package user

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/user/user"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type userCommander struct {
	bot           *tgbotapi.BotAPI
	userCommander Commander
}

func NewUserCommander(bot *tgbotapi.BotAPI) *userCommander {
	return &userCommander{
		bot:           bot,
		userCommander: user.NewUserUserCommander(bot),
	}
}

func (c *userCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "user":
		c.userCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("UserCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *userCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "user":
		c.userCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("UserCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
