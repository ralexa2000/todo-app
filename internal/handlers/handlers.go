package handlers

import (
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/commander"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/storage"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var BadArgument = errors.New("bad argument")

const (
	helpCmd    = "help"
	helpHelp   = "/help - list all commands"
	listCmd    = "list"
	listHelp   = "/list - list all current tasks"
	addCmd     = "add"
	addHelp    = "/add <due_date> <task> - add a new task, due date: YYYY-MM-DD"
	updateCmd  = "update"
	updateHelp = "/update <task_id> <due_date> <task> - update task with id, due date: YYYY-MM-DD"
	deleteCmd  = "delete"
	deleteHelp = "/delete <task_id> - delete task with id"
)

var (
	reAdd    = regexp.MustCompile(`^/add (\d{4}-\d{2}-\d{2}) (.+)$`)
	reUpdate = regexp.MustCompile(`^/update (\d+) (\d{4}-\d{2}-\d{2}) (.+)$`)
	reDelete = regexp.MustCompile(`^/delete (\d+)$`)
)

func AddHandlers(c *commander.Commander) {
	c.RegisterHandler(helpCmd, helpFunc)
	c.RegisterHandler(listCmd, listFunc)
	c.RegisterHandler(addCmd, addFunc)
	c.RegisterHandler(updateCmd, updateFunc)
	c.RegisterHandler(deleteCmd, deleteFunc)
}

func helpFunc(_ ...string) string {
	return strings.Join([]string{
		helpHelp,
		listHelp,
		addHelp,
		updateHelp,
		deleteHelp,
	}, "\n")
}

func listFunc(args ...string) string {
	userName := args[0]
	data := storage.List(userName)
	res := make([]string, 0, len(data))
	for _, t := range data {
		res = append(res, t.String())
	}
	outString := strings.Join(res, "\n")
	if outString == "" {
		outString = "no tasks, add some with /add command"
	}
	return outString
}

func addFunc(args ...string) string {
	userName, inputString := args[0], args[1]
	matched := reAdd.FindStringSubmatch(inputString)
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

func updateFunc(args ...string) string {
	userName, inputString := args[0], args[1]
	matched := reUpdate.FindStringSubmatch(inputString)
	log.Printf("%q\n", matched)
	if len(matched) != 4 {
		return BadArgument.Error()
	}
	id, _ := strconv.ParseUint(matched[1], 10, 64)
	t, err := storage.Get(userName, uint(id))
	if err != nil {
		return err.Error()
	}
	if err = t.SetTask(matched[3]); err != nil {
		return err.Error()
	}
	if err = t.SetDueDate(matched[2]); err != nil {
		return err.Error()
	}
	if err = storage.Update(t); err != nil {
		return err.Error()
	}
	return "task updated"
}

func deleteFunc(args ...string) string {
	userName, inputString := args[0], args[1]
	matched := reDelete.FindStringSubmatch(inputString)
	log.Printf("%q\n", matched)
	if len(matched) != 2 {
		return BadArgument.Error()
	}
	id, _ := strconv.ParseUint(matched[1], 10, 64)
	t, err := storage.Get(userName, uint(id))
	if err != nil {
		return err.Error()
	}
	if err = storage.Delete(t); err != nil {
		return err.Error()
	}
	return "task deleted"
}
