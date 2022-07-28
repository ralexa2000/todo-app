package list

import (
	commandPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command"
	taskPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task"
	"strings"
)

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
	return ""
}

func (c *command) Description() string {
	return "list all current tasks"
}

func (c *command) Process(userName string, _ string) string {
	data := c.task.List(userName)
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
