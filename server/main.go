package main

import (
	"context"
	"log"
	"net"

	pb "avei-grpc/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	PORT = ":50051"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

func (s* TodoServer) CreateTodo (ctx context.Context, in *pb.NewTodo) (*pb.Todo, error) {
	log.Printf("Received: %v", in.GetName())

	todo := &pb.Todo{
		Name :  in.GetName(),
		Description : in.GetDescription(),
		Done: false,
		Id: uuid.New().String(),
	}

	return todo, nil
}

func main() {
	lis, err := net.Listen("tcp", PORT) 

	if err != nil {
		log.Fatalf("failed to Connect: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTodoServiceServer(s, &TodoServer{})

	log.Printf("Server started on port %v", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}