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
var NoAccess = errors.New("no access to task")

func AddHandlers(c *commander.Commander) {
	c.RegisterHandler("list", listFunc)
	c.RegisterHandler("add", addFunc)
	c.RegisterHandler("update", updateFunc)
	c.RegisterHandler("delete", deleteFunc)
}

func listFunc(userName, _ string) string {
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
	if err = storage.Update(t); err != nil {
		return err.Error()
	}
	return "task updated"
}

func deleteFunc(userName, inputString string) string {
	re := regexp.MustCompile(`^/delete (\d+)$`)
	matched := re.FindStringSubmatch(inputString)
	log.Printf("%q\n", matched)
	if len(matched) != 2 {
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
	if err = storage.Delete(t); err != nil {
		return err.Error()
	}
	return "task deleted"
}
