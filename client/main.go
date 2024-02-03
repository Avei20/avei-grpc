package main

import (
	"context"
	"log"
	"time"

	pb "avei-grpc/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type TodoTask struct {
	Name 	  string
	Description string
	Done	bool
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if (err != nil) {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	todos := []TodoTask{
		{Name: "First Task", Description: "This is the first task", Done: false},
		{Name: "Second Task", Description: "This is the second task", Done: false},
		{Name: "Third Task", Description: "This is the third task", Done: false},
	}

	for _, todo := range todos {
		res, err := c.CreateTodo(ctx, &pb.NewTodo{Name: todo.Name, Description: todo.Description})

		if err != nil {
			log.Fatalf("could not create: %v", err)
		}

		log.Printf(`
			ID: %s,
			Name: %s,
			Description: %s,
			Done : %v
		`, res.GetId(), res.GetName(), res.GetDescription(), res.GetDone())
	}
}