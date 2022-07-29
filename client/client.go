package main

import (
	"context"
	pb "gitlab.ozon.dev/ralexa2000/todo-bot/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewAdminClient(conn)

	ctx := context.Background()
	response, err := client.TaskList(ctx, &pb.TaskListRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("response: [%v]", response)
}
