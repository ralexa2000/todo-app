package storage

import (
	"github.com/pkg/errors"
	"sort"
	"strconv"
)

var data map[string]map[uint]*Task // userName -> taskId -> task

var TaskExists = errors.New("task exists")
var TaskNotExists = errors.New("task does not exist")

func init() {
	data = make(map[string]map[uint]*Task)
}

func List(userName string) []*Task {
	tasks, ok := data[userName]
	if !ok {
		return []*Task{}
	}
	res := make([]*Task, 0, len(tasks))
	for _, t := range tasks {
		res = append(res, t)
	}
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].dueDate < res[j].dueDate ||
			(res[i].dueDate == res[j].dueDate && res[i].id < res[j].id)
	})
	return res
}

func Get(userName string, id uint) (*Task, error) {
	if tasks, ok := data[userName]; ok {
		if _, ok := tasks[id]; ok {
			return tasks[id], nil
		}
	}
	return nil, errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(id), 10))
}

func Add(t *Task) error {
	tasks, ok := data[t.GetUser()]
	if !ok {
		data[t.GetUser()] = make(map[uint]*Task)
		tasks = data[t.GetUser()]
	}
	if _, ok := tasks[t.GetId()]; ok {
		return errors.Wrap(TaskExists, strconv.FormatUint(uint64(t.GetId()), 10))
	}
	tasks[t.GetId()] = t
	return nil
}

func Update(t *Task) error {
	if tasks, ok := data[t.GetUser()]; ok {
		if _, ok := tasks[t.GetId()]; ok {
			tasks[t.GetId()] = t
			return nil
		}
	}
	return errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(t.GetId()), 10))
}

func Delete(t *Task) error {
	if tasks, ok := data[t.GetUser()]; ok {
		if _, ok := tasks[t.GetId()]; ok {
			delete(tasks, t.GetId())
			return nil
		}
	}
	return errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(t.GetId()), 10))
}
