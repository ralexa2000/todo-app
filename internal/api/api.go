package api

import (
	"context"
	"github.com/pkg/errors"
	taskPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task"
	cacheLocalPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task/cache/local"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task/models"
	pb "gitlab.ozon.dev/ralexa2000/todo-bot/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var lastIds = make(map[string]uint)

const layoutISO = "2006-01-02"

func New(task taskPkg.Interface) pb.AdminServer {
	return &implementation{
		task: task,
	}
}

type implementation struct {
	pb.UnimplementedAdminServer
	task taskPkg.Interface
}

func (i *implementation) TaskCreate(_ context.Context, in *pb.TaskCreateRequest) (*pb.TaskCreateResponse, error) {
	// parse dueDate
	dueDateParsed, err := time.Parse(layoutISO, in.GetDueDate())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// update lastIds for current user
	if _, ok := lastIds[in.GetUser()]; !ok {
		lastIds[in.GetUser()] = 1
	} else {
		lastIds[in.GetUser()]++
	}

	// create task
	if err := i.task.Create(models.Task{
		Id:      lastIds[in.GetUser()],
		User:    in.GetUser(),
		Task:    in.GetTask(),
		DueDate: dueDateParsed,
	}); err != nil {
		if errors.Is(err, taskPkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.TaskCreateResponse{}, nil
}

func (i *implementation) TaskList(_ context.Context, in *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	tasks := i.task.List(in.GetUser())
	result := make([]*pb.TaskListResponse_Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, &pb.TaskListResponse_Task{
			TaskId:  uint64(task.Id),
			Task:    task.Task,
			DueDate: task.DueDate.Format(layoutISO),
		})
	}
	return &pb.TaskListResponse{
		Tasks: result,
	}, nil
}

func (i *implementation) TaskUpdate(_ context.Context, in *pb.TaskUpdateRequest) (*pb.TaskUpdateResponse, error) {
	// parse dueDate
	dueDateParsed, err := time.Parse(layoutISO, in.GetDueDate())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// find task by its id
	task, err := i.task.Get(in.GetUser(), uint(in.GetTaskId()))
	if err != nil {
		if errors.Is(err, cacheLocalPkg.ErrTaskNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// update task
	task.Task = in.GetTask()
	task.DueDate = dueDateParsed
	if err := i.task.Update(task); err != nil {
		if errors.Is(err, taskPkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.TaskUpdateResponse{}, nil
}

func (i *implementation) TaskDelete(_ context.Context, in *pb.TaskDeleteRequest) (*pb.TaskDeleteResponse, error) {
	// find task by its id
	task, err := i.task.Get(in.GetUser(), uint(in.GetTaskId()))
	if err != nil {
		if errors.Is(err, cacheLocalPkg.ErrTaskNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// delete task
	if err := i.task.Delete(task); err != nil {
		if errors.Is(err, taskPkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.TaskDeleteResponse{}, nil
}
