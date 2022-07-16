package storage

import "fmt"

var lastId = uint(0)

type Task struct {
	id      uint
	user    string
	task    string
	dueDate string
}

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
	lastId++
	t.id = lastId
	return &t, nil
}

func (t *Task) String() string {
	return fmt.Sprintf("[id %d] [till %v] %v", t.id, t.dueDate, t.task)
}

func (t *Task) SetUser(user string) error {
	if len(user) < 3 || len(user) > 10 {
		return fmt.Errorf("bad user name length <%v>", user)
	}
	t.user = user
	return nil
}

func (t *Task) SetTask(task string) error {
	if len(task) == 0 || len(task) > 100 {
		return fmt.Errorf("bad task length <%v>", task)
	}
	t.task = task
	return nil
}

func (t *Task) SetDueDate(dueDate string) error {
	if len(dueDate) != 10 {
		return fmt.Errorf("bad dueDate length <%v>", dueDate)
	}
	t.dueDate = dueDate
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
	return t.dueDate
}
