package cache

import (
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task/models"
)

type Interface interface {
	Create(task models.Task) error
	Update(task models.Task) error
	List(userName string) []models.Task
	Delete(task models.Task) error
	String(task models.Task) string
}
