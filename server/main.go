package main

import (
	"context"
	pb "gatewaytest/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {}

func (s *server) Post(ctx context.Context, in *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{Message: "hello post "+in.Name}, nil
}
func (s *server) Get(ctx context.Context, in *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{Message: "hello get "+in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTestServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}