package main

import (
	apiPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/api"
	pb "gitlab.ozon.dev/ralexa2000/todo-bot/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func runGRPCServer() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, apiPkg.New())

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
