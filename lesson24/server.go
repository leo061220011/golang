package main

import (
	"context"
	"fmt"
	pb "lesson24/proto"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + req.GetName()}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	fmt.Println("Server started... on port 50051")
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
		return
	}
}
