package commander

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ralexa2000/todo-bot/config"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/storage"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	listCmd   = "list"
	addCmd    = "add"
	updateCmd = "update"
)

type Commander struct {
	bot *tgbotapi.BotAPI
}

var UnknownCommand = errors.New("unknown command")
var BadArgument = errors.New("bad argument")
var NoAccess = errors.New("no access to task")

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
			case addCmd:
				msg.Text = addFunc(userName, update.Message.Text)
			case updateCmd:
				msg.Text = updateFunc(userName, update.Message.Text)
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

func addFunc(userName, inputString string) string {
	re := regexp.MustCompile(`^/add (\d{4}-\d{2}-\d{2}) (.+)$`)
	matched := re.FindStringSubmatch(inputString)
	log.Printf("%q\n", matched)
	if len(matched) != 3 {
		return BadArgument.Error()
	}
	t, err := storage.NewTask(userName, matched[2], matched[1])
	if err != nil {
		return err.Error()
	}
	err = storage.Add(t)
	if err != nil {
		return err.Error()
	}
	return "task added"
}

func updateFunc(userName, inputString string) string {
	re := regexp.MustCompile(`^/update (\d+) (\d{4}-\d{2}-\d{2}) (.+)$`)
	matched := re.FindStringSubmatch(inputString)
	log.Printf("%q\n", matched)
	if len(matched) != 4 {
		return BadArgument.Error()
	}
	id, _ := strconv.ParseUint(matched[1], 10, 64)
	t, err := storage.GetById(uint(id))
	if err != nil {
		return err.Error()
	}
	if t.GetUser() != userName {
		return NoAccess.Error()
	}
	if err = t.SetTask(matched[3]); err != nil {
		return err.Error()
	}
	if err = t.SetDueDate(matched[2]); err != nil {
		return err.Error()
	}
	err = storage.Update(t)
	return "task updated"
}
