package storage

import (
	"github.com/pkg/errors"
	"log"
	"strconv"
)

var data map[uint]*Task

var TaskExists = errors.New("task exists")
var TaskNotExists = errors.New("task does not exist")

func init() {
	data = make(map[uint]*Task)
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
	res := make([]*Task, 0, len(data))
	for _, t := range data {
		if t.user == userName {
			res = append(res, t)
		}
	}
	return res
}

func Add(t *Task) error {
	if _, ok := data[t.GetId()]; ok {
		return errors.Wrap(TaskExists, strconv.FormatUint(uint64(t.GetId()), 10))
	}
	data[t.GetId()] = t
	return nil
}

func Update(t *Task) error {
	if _, ok := data[t.GetId()]; !ok {
		return errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(t.GetId()), 10))
	}
	data[t.GetId()] = t
	return nil
}

func Delete(t *Task) error {
	if _, ok := data[t.GetId()]; !ok {
		return errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(t.GetId()), 10))
	}
	delete(data, t.GetId())
	return nil
}
