package delete

import (
	"fmt"
	"github.com/pkg/errors"
	commandPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command"
	taskPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/repository"
	"log"
	"regexp"
	"strconv"
)

var (
	regexpDelete = regexp.MustCompile(`^/delete (\d+)$`)
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
	return "delete"
}

func (c *command) Arguments() string {
	return "<task_id>"
}

func (c *command) Description() string {
	return "delete task with id"
}

func (c *command) Process(userName string, inputString string) string {
	// parse inputString into arguments
	matched := regexpDelete.FindStringSubmatch(inputString)
	log.Printf("%q\n", matched)
	if len(matched) != 2 {
		return "invalid args"
	}
	taskId := matched[1]

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

	// delete task
	if err = c.task.Delete(task); err != nil {
		if errors.Is(err, taskPkg.ErrValidation) {
			return "invalid args"
		}
		log.Printf("INTERNAL ERROR: %s", err.Error())
		return "internal error"
	}

	return "task deleted"
}
