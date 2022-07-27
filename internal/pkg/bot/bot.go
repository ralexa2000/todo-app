package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/config"
	commandPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command"
	"log"
)

var unknownCommandError = "unknown command"

type Interface interface {
	RegisterHandler(cmd commandPkg.Interface)
	Run() error
}

func MustNew() Interface {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "init bot"))
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return &commander{
		bot:   bot,
		route: make(map[string]commandPkg.Interface),
	}
}

type commander struct {
	bot   *tgbotapi.BotAPI
	route map[string]commandPkg.Interface
}

func (c *commander) RegisterHandler(cmd commandPkg.Interface) {
	c.route[cmd.Name()] = cmd
}

func (c *commander) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		userName := update.Message.From.UserName
		userText := update.Message.Text
		if cmdName := update.Message.Command(); cmdName != "" {
			if cmd, ok := c.route[cmdName]; ok {
				msg.Text = cmd.Process(userName, userText)
			} else {
				msg.Text = fmt.Sprintf("%s\n\n%s", unknownCommandError, c.route["help"].Process("", ""))
			}
		} else {
			log.Printf("[%s] %s", userName, userText)
			msg.Text = fmt.Sprintf("you sent me <%s>\n\n%s", userText, c.route["help"].Process("", ""))
		}
		_, err := c.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send message to bot")
		}
	}
	return nil
}
