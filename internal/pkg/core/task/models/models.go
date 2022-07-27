package models

import "time"

type Task struct {
	Id      uint
	User    string
	Task    string
	DueDate time.Time
}
