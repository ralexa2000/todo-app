package api

import (
	"context"
	pb "gitlab.ozon.dev/ralexa2000/todo-bot/pkg/api"
)

func New() pb.AdminServer {
	return &implementation{}
}

type implementation struct {
	pb.UnimplementedAdminServer
}

func (i *implementation) TaskCreate(ctx context.Context, in *pb.TaskCreateRequest) (*pb.TaskCreateResponse, error) {
	return nil, nil
}

func (i *implementation) TaskList(ctx context.Context, in *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	return &pb.TaskListResponse{
		Tasks: []*pb.TaskListResponse_Task{{
			TaskId:  1,
			User:    "mock",
			Task:    "mock",
			DueDate: "mock",
		}},
	}, nil
}

func (i *implementation) TaskUpdate(ctx context.Context, in *pb.TaskUpdateRequest) (*pb.TaskUpdateResponse, error) {
	return nil, nil
}

func (i *implementation) TaskDelete(ctx context.Context, in *pb.TaskDeleteRequest) (*pb.TaskDeleteResponse, error) {
	return nil, nil
}
