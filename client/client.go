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
	createResponse, err := client.TaskCreate(ctx, &pb.TaskCreateRequest{
		User:    "ralexa2000",
		Task:    "clean my room",
		DueDate: "2022-07-31",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("createResponse: [%v]", createResponse)

	listResponse, err := client.TaskList(ctx, &pb.TaskListRequest{
		User: "ralexa2000",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listResponse: [%v]", listResponse)

	updateResponse, err := client.TaskUpdate(ctx, &pb.TaskUpdateRequest{
		TaskId:  1,
		User:    "ralexa2000",
		Task:    "clean my room",
		DueDate: "2022-08-01",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("updateResponse: [%v]", updateResponse)

	listResponse, err = client.TaskList(ctx, &pb.TaskListRequest{
		User: "ralexa2000",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listResponse: [%v]", listResponse)

	deleteResponse, err := client.TaskDelete(ctx, &pb.TaskDeleteRequest{
		TaskId: 1,
		User:   "ralexa2000",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("deleteResponse: [%v]", deleteResponse)

	listResponse, err = client.TaskList(ctx, &pb.TaskListRequest{
		User: "ralexa2000",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listResponse: [%v]", listResponse)
}
