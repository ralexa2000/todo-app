package list

import (
	commandPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command"
	taskPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var regexpList = regexp.MustCompile(`^/list (\d+) (\d+)$`)

func New(task taskPkg.Interface) commandPkg.Interface {
	return &command{
		task: task,
	}
}

type command struct {
	task taskPkg.Interface
}

func (c *command) Name() string {
	return "list"
}

func (c *command) Arguments() string {
	return "<limit> <offset>"
}

func (c *command) Description() string {
	return "list all current tasks, use limit and offset for pagination (required)"
}

func (c *command) Process(userName string, inputString string) string {
	// parse inputString into arguments
	matched := regexpList.FindStringSubmatch(inputString)
	log.Printf("%q\n", matched)
	if len(matched) != 3 {
		return "invalid args"
	}
	limitStr, offsetStr := matched[1], matched[2]

	// parse limit and offset
	limit, _ := strconv.ParseUint(limitStr, 10, 32)
	offset, _ := strconv.ParseUint(offsetStr, 10, 32)

	data := c.task.List(userName, uint32(limit), uint32(offset))
	res := make([]string, 0, len(data))
	for _, task := range data {
		res = append(res, c.task.String(task))
	}
	outString := strings.Join(res, "\n")
	if outString == "" {
		outString = "no tasks, add some with /add command"
	}
	return outString
}
