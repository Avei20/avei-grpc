package main

import (
	"context"
	"log"
	"net"

	// pb "buf.build/gen/go/avei/proto/protocolbuffers/go/todo/v1"
	pb "buf.build/gen/go/avei/proto/grpc/go/todo/v1/todov1grpc"
	// "buf.build/gen/go/avei/proto/protocolbuffers/go"

	v1 "buf.build/gen/go/avei/proto/protocolbuffers/go/todo/v1"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	PORT = ":50051"
)

// moduleOwner = avei 
// moduleName = proto 
// pluginOwner = protocolbuffers 
// pluginName = go 





type TodoServer struct {
	pb.TodoServiceServer
}

func (s* TodoServer) CreateTodo (ctx context.Context, in *v1.CreateTodoRequest) (*v1.CreateTodoResponse, error) {
	log.Printf("Received: %v", in.GetName())

	todo := &v1.CreateTodoResponse{
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