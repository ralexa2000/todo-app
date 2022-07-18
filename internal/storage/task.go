package storage

import (
	"fmt"
	"time"
)

var lastIds = make(map[string]uint)

type Task struct {
	id      uint
	user    string
	task    string
	dueDate time.Time
}

const layoutISO = "2006-01-02"

func NewTask(user, task, dueDate string) (*Task, error) {
	t := Task{}
	if err := t.SetUser(user); err != nil {
		return nil, err
	}
	if err := t.SetTask(task); err != nil {
		return nil, err
	}
	if err := t.SetDueDate(dueDate); err != nil {
		return nil, err
	}
	if _, ok := lastIds[user]; !ok {
		lastIds[user] = 1
	} else {
		lastIds[user]++
	}
	t.id = lastIds[user]
	return &t, nil
}

func (t *Task) String() string {
	return fmt.Sprintf("[id %d] [till %s] %s", t.GetId(), t.GetDueDate(), t.GetTask())
}

func (t *Task) SetUser(user string) error {
	if len(user) < 3 || len(user) > 10 {
		return fmt.Errorf("bad user name length <%s>", user)
	}
	t.user = user
	return nil
}

func (t *Task) SetTask(task string) error {
	if len(task) == 0 || len(task) > 100 {
		return fmt.Errorf("bad task length <%s>", task)
	}
	t.task = task
	return nil
}

func (t *Task) SetDueDate(dueDate string) error {
	dueDateParsed, err := time.Parse(layoutISO, dueDate)
	if err != nil {
		return fmt.Errorf("bad dueDate <%s>", dueDate)
	}
	t.dueDate = dueDateParsed
	return nil
}

func (t *Task) GetId() uint {
	return t.id
}

func (t *Task) GetUser() string {
	return t.user
}

func (t *Task) GetTask() string {
	return t.task
}

func (t *Task) GetDueDate() string {
	return t.dueDate.Format(layoutISO)
}
