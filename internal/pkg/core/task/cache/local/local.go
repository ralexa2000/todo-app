package local

import (
	"fmt"
	"github.com/pkg/errors"
	cachePkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task/cache"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task/models"
	"sort"
	"sync"
)

const poolSize = 10

var (
	ErrTaskExists    = errors.New("task exists")
	ErrTaskNotExists = errors.New("task does not exist")
)

const layoutISO = "2006-01-02"

func New() cachePkg.Interface {
	return &cache{
		mu:     sync.RWMutex{},
		data:   map[string]map[uint]models.Task{},
		poolCh: make(chan struct{}, poolSize),
	}
}

type cache struct {
	mu     sync.RWMutex
	data   map[string]map[uint]models.Task
	poolCh chan struct{}
}

func (c *cache) Create(task models.Task) error {
	c.mu.Lock()
	c.poolCh <- struct{}{}
	defer func() {
		defer c.mu.Unlock()
		<-c.poolCh
	}()

	tasks, ok := c.data[task.User]
	if !ok {
		c.data[task.User] = make(map[uint]models.Task)
		tasks = c.data[task.User]
	}
	if _, ok := tasks[task.Id]; ok {
		return errors.Wrapf(ErrTaskExists, "task_id: [%d]", task.Id)
	}
	tasks[task.Id] = task
	return nil
}

func (c *cache) Update(task models.Task) error {
	c.mu.Lock()
	c.poolCh <- struct{}{}
	defer func() {
		defer c.mu.Unlock()
		<-c.poolCh
	}()

	if tasks, ok := c.data[task.User]; ok {
		if _, ok := tasks[task.Id]; ok {
			tasks[task.Id] = task
			return nil
		}
	}
	return errors.Wrapf(ErrTaskNotExists, "task_id: [%d]", task.Id)
}

func (c *cache) List(userName string) []models.Task {
	c.mu.RLock()
	c.poolCh <- struct{}{}
	defer func() {
		defer c.mu.RUnlock()
		<-c.poolCh
	}()

	tasks, ok := c.data[userName]
	if !ok {
		return []models.Task{}
	}
	result := make([]models.Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, task)
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].DueDate.Before(result[j].DueDate) ||
			(result[i].DueDate.Equal(result[j].DueDate) && result[i].Id < result[j].Id)
	})
	return result
}

func (c *cache) Delete(task models.Task) error {
	c.mu.Lock()
	c.poolCh <- struct{}{}
	defer func() {
		defer c.mu.Unlock()
		<-c.poolCh
	}()

	if tasks, ok := c.data[task.User]; ok {
		if _, ok := tasks[task.Id]; ok {
			delete(tasks, task.Id)
			return nil
		}
	}
	return errors.Wrapf(ErrTaskNotExists, "task_id: [%d]", task.Id)
}

func (c *cache) Get(userName string, taskId uint) (models.Task, error) {
	c.mu.RLock()
	c.poolCh <- struct{}{}
	defer func() {
		defer c.mu.RUnlock()
		<-c.poolCh
	}()

	if tasks, ok := c.data[userName]; ok {
		if _, ok := tasks[taskId]; ok {
			return tasks[taskId], nil
		}
	}
	return models.Task{}, errors.Wrapf(ErrTaskNotExists, "task_id: [%d]", taskId)
}

func (c *cache) String(task models.Task) string {
	return fmt.Sprintf("[id %d] [till %s] %s", task.Id, task.DueDate.Format(layoutISO), task.Task)
}
