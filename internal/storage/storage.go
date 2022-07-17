package storage

import (
	"github.com/pkg/errors"
	"log"
	"strconv"
)

var data map[string]map[uint]*Task // userName -> taskId -> task

var TaskExists = errors.New("task exists")
var TaskNotExists = errors.New("task does not exist")

func init() {
	data = make(map[string]map[uint]*Task)
	t1, _ := NewTask("ralexa2000", "create telegram bot in go", "2022-07-24")
	if err := Add(t1); err != nil {
		log.Panic(err)
	}
	t2, _ := NewTask("other_user", "create telegram bot in go", "2022-07-24")
	if err := Add(t2); err != nil {
		log.Panic(err)
	}
}

func List(userName string) []*Task {
	if _, ok := data[userName]; !ok {
		return []*Task{}
	}
	res := make([]*Task, 0, len(data[userName]))
	for _, t := range data[userName] {
		if t.user == userName {
			res = append(res, t)
		}
	}
	return res
}

func Get(userName string, id uint) (*Task, error) {
	tasks, ok := data[userName]
	if !ok {
		return nil, errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(id), 10))
	}
	if _, ok := tasks[id]; !ok {
		return nil, errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(id), 10))
	}
	return tasks[id], nil
}

func Add(t *Task) error {
	tasks, ok := data[t.user]
	if !ok {
		data[t.user] = make(map[uint]*Task)
		tasks = data[t.user]
	}
	if _, ok := tasks[t.GetId()]; ok {
		return errors.Wrap(TaskExists, strconv.FormatUint(uint64(t.GetId()), 10))
	}
	tasks[t.GetId()] = t
	return nil
}

func Update(t *Task) error {
	tasks, ok := data[t.user]
	if !ok {
		return errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(t.GetId()), 10))
	}
	if _, ok := tasks[t.GetId()]; !ok {
		return errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(t.GetId()), 10))
	}
	tasks[t.GetId()] = t
	return nil
}

func Delete(t *Task) error {
	tasks, ok := data[t.user]
	if !ok {
		return errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(t.GetId()), 10))
	}
	if _, ok := tasks[t.GetId()]; !ok {
		return errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(t.GetId()), 10))
	}
	delete(tasks, t.GetId())
	return nil
}
