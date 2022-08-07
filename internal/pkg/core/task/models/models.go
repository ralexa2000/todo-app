package models

import "time"

type Task struct {
	Id      uint      `db:"task_id"`
	User    string    `db:"username"`
	Task    string    `db:"task"`
	DueDate time.Time `db:"due_date"`
}
