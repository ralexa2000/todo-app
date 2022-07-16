package commander

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ralexa2000/todo-bot/config"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/storage"
	"log"
	"strings"
)

const (
	listCmd = "list"
)

type Commander struct {
	bot *tgbotapi.BotAPI
}

var UnknownCommand = errors.New("unknown command")

func Init() (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		return nil, errors.Wrap(err, "init bot")
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return &Commander{bot: bot}, nil
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
			switch cmd {
			case listCmd:
				msg.Text = listFunc(userName)
			default:
				msg.Text = UnknownCommand.Error()
			}
		} else {
			log.Printf("[%s] %s", userName, update.Message.Text)
			msg.Text = fmt.Sprintf("you sent me <%v>", update.Message.Text)
		}
		_, err := c.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send message to bot")
		}
	}
	return nil
}

func listFunc(userName string) string {
	data := storage.List(userName)
	res := make([]string, 0, len(data))
	for _, t := range data {
		res = append(res, t.String())
	}
	return strings.Join(res, "\n")
}
