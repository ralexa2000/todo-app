package commander

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ralexa2000/todo-bot/config"
	"log"
)

type Commander struct {
	bot *tgbotapi.BotAPI
}

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
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.Text = fmt.Sprintf("you sent me <%v>", update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			_, err := c.bot.Send(msg)
			if err != nil {
				return errors.Wrap(err, "send message to bot")
			}
		}
	}
	return nil
}
