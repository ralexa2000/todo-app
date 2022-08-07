package api

import (
	"context"
	"github.com/pkg/errors"
	taskPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task/models"
	"gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/repository"
	pb "gitlab.ozon.dev/ralexa2000/todo-bot/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sync"
	"time"
)

var (
	lastIds   = make(map[string]uint)
	muLastIds = sync.RWMutex{}
)

const (
	layoutISO        = "2006-01-02"
	listLimitDefault = 10
)

func New(task taskPkg.Interface) pb.AdminServer {
	return &implementation{
		task: task,
	}
}

type implementation struct {
	pb.UnimplementedAdminServer
	task taskPkg.Interface
}

func updateLastIds(user string) {
	muLastIds.Lock()
	defer muLastIds.Unlock()

	if _, ok := lastIds[user]; !ok {
		lastIds[user] = 1
	} else {
		lastIds[user]++
	}
}

func getLastIds(user string) uint {
	muLastIds.RLock()
	defer muLastIds.RUnlock()

	return lastIds[user]
}

func (i *implementation) TaskCreate(_ context.Context, in *pb.TaskCreateRequest) (*pb.TaskCreateResponse, error) {
	// parse dueDate
	dueDateParsed, err := time.Parse(layoutISO, in.GetDueDate())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// update lastIds for current user
	updateLastIds(in.GetUser())

	// create task
	if err := i.task.Create(models.Task{
		Id:      getLastIds(in.GetUser()),
		User:    in.GetUser(),
		Task:    in.GetTask(),
		DueDate: dueDateParsed,
	}); err != nil {
		if errors.Is(err, taskPkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		log.Println("INTERNAL ERROR:", err.Error())
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &pb.TaskCreateResponse{}, nil
}

func (i *implementation) TaskGet(_ context.Context, in *pb.TaskGetRequest) (*pb.TaskGetResponse, error) {
	task, err := i.task.Get(in.GetUser(), uint(in.GetTaskId()))
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		log.Println("INTERNAL ERROR:", err.Error())
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &pb.TaskGetResponse{
		TaskId:  uint64(task.Id),
		User:    task.User,
		Task:    task.Task,
		DueDate: task.DueDate.Format(layoutISO),
	}, nil
}

func (i *implementation) TaskList(_ context.Context, in *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	limit := in.GetLimit()
	if limit == 0 {
		limit = listLimitDefault
	}
	tasks := i.task.List(in.GetUser(), limit, in.GetOffset())
	result := make([]*pb.TaskListResponse_Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, &pb.TaskListResponse_Task{
			TaskId:  uint64(task.Id),
			User:    task.User,
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
		if errors.Is(err, repository.ErrTaskNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		log.Println("INTERNAL ERROR:", err.Error())
		return nil, status.Error(codes.Internal, "internal error")
	}

	// update task
	task.Task = in.GetTask()
	task.DueDate = dueDateParsed
	if err := i.task.Update(task); err != nil {
		if errors.Is(err, taskPkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		log.Println("INTERNAL ERROR:", err.Error())
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &pb.TaskUpdateResponse{}, nil
}

func (i *implementation) TaskDelete(_ context.Context, in *pb.TaskDeleteRequest) (*pb.TaskDeleteResponse, error) {
	// find task by its id
	task, err := i.task.Get(in.GetUser(), uint(in.GetTaskId()))
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		log.Println("INTERNAL ERROR:", err.Error())
		return nil, status.Error(codes.Internal, "internal error")
	}

	// delete task
	if err := i.task.Delete(task); err != nil {
		if errors.Is(err, taskPkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		log.Println("INTERNAL ERROR:", err.Error())
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &pb.TaskDeleteResponse{}, nil
}
