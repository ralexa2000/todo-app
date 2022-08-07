package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task/models"
	"log"
)

var (
	ErrTaskNotExists = errors.New("task does not exist")
)

func (r *Repository) TaskGet(ctx context.Context, username string, taskId uint) (models.Task, error) {
	log.Printf("Selecting task id [%d] for username [%s]", taskId, username)
	query := "SELECT username, task_id, task, due_date FROM public.tasks WHERE username = $1 AND task_id = $2"
	var task models.Task
	if err := pgxscan.Get(
		ctx, r.pool, &task, query, username, taskId,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Task{}, errors.Wrapf(ErrTaskNotExists, "user: [%s], task_id: [%d]", username, taskId)
		}
		return models.Task{}, fmt.Errorf("Repository.TaskGet: pgxscan.Get: %w", err)
	}
	return task, nil
}

func (r *Repository) GetNewIdForUser(ctx context.Context, username string) (uint, error) {
	query := "SELECT COALESCE(MAX(task_id), 0) AS max_task_id FROM public.tasks WHERE username = $1"
	row := r.pool.QueryRow(ctx, query, username)
	var maxTaskId uint
	if err := row.Scan(&maxTaskId); err != nil {
		return 0, fmt.Errorf("Repository.GetNewIdForUser: Scan: %w", err)
	}
	maxTaskId += 1
	return maxTaskId, nil
}

func (r *Repository) TaskCreate(ctx context.Context, task *models.Task) error {
	log.Printf(
		"Creating new task, username [%s], task_id [%d], task [%s], due_date [%s]",
		task.User, task.Id, task.Task, task.DueDate,
	)
	query := "INSERT INTO public.tasks (username, task_id, task, due_date) VALUES ($1, $2, $3, $4)"
	_, err := r.pool.Exec(ctx, query, task.User, task.Id, task.Task, task.DueDate)
	if err != nil {
		return fmt.Errorf("Repository.TaskCreate: Exec: %w", err)
	}
	return nil
}

func (r *Repository) TaskUpdate(ctx context.Context, task *models.Task) error {
	log.Printf(
		"Updating task, username [%s], task_id [%d], task [%s], due_date [%s]",
		task.User, task.Id, task.Task, task.DueDate,
	)
	query := "UPDATE public.tasks SET task = $3, due_date = $4 WHERE username = $1 AND task_id = $2"
	_, err := r.pool.Exec(ctx, query, task.User, task.Id, task.Task, task.DueDate)
	if err != nil {
		return fmt.Errorf("Repository.TaskUpdate: Exec: %w", err)
	}
	return nil
}

func (r *Repository) TaskDelete(ctx context.Context, task *models.Task) error {
	log.Printf(
		"Deleting task, username [%s], task_id [%d], task [%s], due_date [%s]",
		task.User, task.Id, task.Task, task.DueDate,
	)
	query := "DELETE FROM public.tasks WHERE username = $1 AND task_id = $2"
	_, err := r.pool.Exec(ctx, query, task.User, task.Id)
	if err != nil {
		return fmt.Errorf("Repository.TaskDelete: Exec: %w", err)
	}
	return nil
}

func (r *Repository) TaskList(ctx context.Context, username string, limit, offset uint32) []models.Task {
	log.Printf("Selecting tasks for username [%s]", username)
	query := fmt.Sprintln(
		"SELECT username, task_id, task, due_date FROM public.tasks WHERE username = $1",
		"ORDER BY due_date LIMIT $2 OFFSET $3",
	)
	var tasks []models.Task
	if err := pgxscan.Select(ctx, r.pool, &tasks, query, username, limit, offset); err != nil {
		log.Printf("INTERNAL ERROR: Repository.TaskList: username [%s], err: %s", username, err.Error())
		return []models.Task{}
	}
	return tasks
}
