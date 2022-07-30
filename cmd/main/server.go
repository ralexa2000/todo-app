package main

import (
	apiPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/api"
	taskPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/core/task"
	pb "gitlab.ozon.dev/ralexa2000/todo-bot/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func runGRPCServer() {
	var task taskPkg.Interface
	{
		task = taskPkg.New()
	}

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, apiPkg.New(task))

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
