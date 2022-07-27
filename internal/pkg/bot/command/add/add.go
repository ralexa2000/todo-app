package add

import (
	"fmt"
	"github.com/pkg/errors"
	commandPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command"
	taskPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task/models"
	"log"
	"regexp"
	"time"
)

var reAdd = regexp.MustCompile(`^/add (\d{4}-\d{2}-\d{2}) (.+)$`)
var lastIds = make(map[string]uint)

const layoutISO = "2006-01-02"

func New(task taskPkg.Interface) commandPkg.Interface {
	return &command{
		task: task,
	}
}

type command struct {
	task taskPkg.Interface
}

func (c *command) Name() string {
	return "add"
}

func (c *command) Arguments() string {
	return "<due_date> <task>"
}

func (c *command) Description() string {
	return "add a new task, due date: YYYY-MM-DD"
}

func (c *command) Process(userName string, inputString string) string {
	// parse inputString into arguments
	matched := reAdd.FindStringSubmatch(inputString)
	log.Printf("%q\n", matched)
	if len(matched) != 3 {
		return "invalid args"
	}
	dueDate, task := matched[1], matched[2]

	// parse dueDate
	dueDateParsed, err := time.Parse(layoutISO, dueDate)
	if err != nil {
		return fmt.Sprintf("bad dueDate <%s>", dueDate)
	}

	// update lastIds for current user
	if _, ok := lastIds[userName]; !ok {
		lastIds[userName] = 1
	} else {
		lastIds[userName]++
	}

	// create new task
	if err := c.task.Create(models.Task{
		Id:      lastIds[userName],
		User:    userName,
		Task:    task,
		DueDate: dueDateParsed,
	}); err != nil {
		if errors.Is(err, taskPkg.ErrValidation) {
			return "invalid args"
		}
		log.Printf("Internal Error: %s", err.Error())
		return "internal error"
	}
	return "task added"
}
