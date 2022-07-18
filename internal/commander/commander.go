package commander

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ralexa2000/todo-bot/config"
	"log"
)

type CmdHandler func(...string) string

type Commander struct {
	bot   *tgbotapi.BotAPI
	route map[string]CmdHandler
}

var unknownCommandError = "unknown command"

func (c *Commander) RegisterHandler(cmd string, handler CmdHandler) {
	c.route[cmd] = handler
}

func Init() (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		return nil, errors.Wrap(err, "init bot")
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return &Commander{
		bot:   bot,
		route: make(map[string]CmdHandler),
	}, nil
}

func (c *Commander) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		userName := update.Message.From.UserName
		if cmd := update.Message.Command(); cmd != "" {
			if handler, ok := c.route[cmd]; ok {
				msg.Text = handler(userName, update.Message.Text)
			} else {
				msg.Text = fmt.Sprintf("%s\n\n%s", unknownCommandError, c.route["help"]())
			}
		} else {
			log.Printf("[%s] %s", userName, update.Message.Text)
			msg.Text = fmt.Sprintf("you sent me <%s>\n\n%s", update.Message.Text, c.route["help"]())
		}
		_, err := c.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send message to bot")
		}
	}
	return nil
}
