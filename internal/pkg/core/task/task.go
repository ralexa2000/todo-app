package task

import (
	"github.com/pkg/errors"
	cachePkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task/cache"
	localCachePkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task/cache/local"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task/models"
)

var (
	ErrValidation = errors.New("invalid data")
)

type Interface interface {
	Create(task models.Task) error
	Update(task models.Task) error
	List(userName string) []models.Task
	Delete(task models.Task) error
}

func New() Interface {
	return &core{
		cache: localCachePkg.New(),
	}
}

type core struct {
	cache cachePkg.Interface
}

func (c *core) Create(task models.Task) error {
	if task.Task == "" {
		return errors.Wrap(ErrValidation, "field: [task] cannot be empty")
	}
	if task.User == "" {
		return errors.Wrap(ErrValidation, "field: [user] cannot be empty")
	}
	return c.cache.Create(task)
}

func (c *core) Update(task models.Task) error {
	if task.Task == "" {
		return errors.Wrap(ErrValidation, "field: [task] cannot be empty")
	}
	if task.User == "" {
		return errors.Wrap(ErrValidation, "field: [user] cannot be empty")
	}
	return c.cache.Update(task)
}

func (c *core) List(userName string) []models.Task {
	return c.cache.List(userName)
}

func (c *core) Delete(task models.Task) error {
	if task.Task == "" {
		return errors.Wrap(ErrValidation, "field: [task] cannot be empty")
	}
	if task.User == "" {
		return errors.Wrap(ErrValidation, "field: [user] cannot be empty")
	}
	return c.cache.Delete(task)
}
