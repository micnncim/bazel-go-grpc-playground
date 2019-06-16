package api

import (
	"context"
	"log"

	pb "github.com/micnncim/bazel-go-grpc-playground/pkg/service"
)

type server struct{}

func NewServer() pb.EchoAPIServer {
	return &server{}
}

func (s *server) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("received request: %s\n", req.Message)
	return &pb.EchoResponse{
		Message: req.Message,
	}, nil
}
