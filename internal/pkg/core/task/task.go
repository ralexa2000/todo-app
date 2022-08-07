package task

import (
	"context"
	"github.com/pkg/errors"
	cachePkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task/cache"
	localCachePkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task/cache/local"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task/models"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/repository"
	"regexp"
)

const regexpUserNameString = `^[A-Za-z0-9_\-\.]+$`

var (
	ErrValidation  = errors.New("invalid data")
	regexpUsername = regexp.MustCompile(regexpUserNameString)
)

type Interface interface {
	Create(task models.Task) error
	Update(task models.Task) error
	List(userName string) []models.Task
	Delete(task models.Task) error
	Get(userName string, taskId uint) (models.Task, error)
	String(task models.Task) string
}

func New(ctx context.Context, repo *repository.Repository) Interface {
	return &core{
		cache: localCachePkg.New(),
		repo:  repo,
		ctx:   ctx,
	}
}

type core struct {
	cache cachePkg.Interface
	repo  *repository.Repository
	ctx   context.Context
}

func (c *core) Create(task models.Task) error {
	if task.Task == "" {
		return errors.Wrap(ErrValidation, "field: [task] cannot be empty")
	}
	if !regexpUsername.MatchString(task.User) {
		return errors.Wrapf(ErrValidation, "field: [user] does not match %s", regexpUserNameString)
	}
	taskId, err := c.repo.GetNewIdForUser(c.ctx, task.User)
	if err != nil {
		return errors.Wrap(err, "task.TaskCreate")
	}
	task.Id = taskId
	if err := c.repo.TaskCreate(c.ctx, &task); err != nil {
		return errors.Wrap(err, "task.TaskCreate")
	}
	return nil
}

func (c *core) Update(task models.Task) error {
	if task.Task == "" {
		return errors.Wrap(ErrValidation, "field: [task] cannot be empty")
	}
	if !regexpUsername.MatchString(task.User) {
		return errors.Wrapf(ErrValidation, "field: [user] does not match %s", regexpUserNameString)
	}

	if err := c.repo.TaskUpdate(c.ctx, &task); err != nil {
		return errors.Wrap(err, "task.TaskUpdate")
	}
	return nil
}

func (c *core) List(userName string) []models.Task {
	return c.repo.TaskList(c.ctx, userName)
}

func (c *core) Delete(task models.Task) error {
	return c.repo.TaskDelete(c.ctx, &task)
}

func (c *core) Get(userName string, taskId uint) (models.Task, error) {
	return c.repo.TaskGet(c.ctx, userName, taskId)
}

func (c *core) String(task models.Task) string {
	return c.cache.String(task)
}
