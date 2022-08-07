package update

import (
	"errors"
	"fmt"
	commandPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command"
	taskPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/repository"
	"log"
	"regexp"
	"strconv"
	"time"
)

var (
	regexpUpdate = regexp.MustCompile(`^/update (\d+) (\d{4}-\d{2}-\d{2}) (.+)$`)
)

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
	return "update"
}

func (c *command) Arguments() string {
	return "<task_id> <due_date:YYYY-MM-DD> <task>"
}

func (c *command) Description() string {
	return "update task with id"
}

func (c *command) Process(userName string, inputString string) string {
	// parse inputString into arguments
	matched := regexpUpdate.FindStringSubmatch(inputString)
	log.Printf("%q\n", matched)
	if len(matched) != 4 {
		return "invalid args"
	}
	taskId, dueDate, taskName := matched[1], matched[2], matched[3]

	// parse dueDate
	dueDateParsed, err := time.Parse(layoutISO, dueDate)
	if err != nil {
		return fmt.Sprintf("bad dueDate <%s>", dueDate)
	}

	// parse taskId
	taskIdParsed, _ := strconv.ParseUint(taskId, 10, 64)

	// find task by its id
	task, err := c.task.Get(userName, uint(taskIdParsed))
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotExists) {
			return fmt.Sprintf("task with id <%d> does not exist", taskIdParsed)
		}
		log.Printf("INTERNAL ERROR: %s", err.Error())
		return "internal error"
	}

	// update task
	task.Task = taskName
	task.DueDate = dueDateParsed
	if err = c.task.Update(task); err != nil {
		if errors.Is(err, taskPkg.ErrValidation) {
			return "invalid args"
		}
		log.Printf("INTERNAL ERROR: %s", err.Error())
		return "internal error"
	}

	return "task updated"
}
