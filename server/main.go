package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/micnncim/bazel-go-grpc-playground/pkg/api"
	pb "github.com/micnncim/bazel-go-grpc-playground/pkg/service"
)

const port = ":8080"

func main() {
	s := grpc.NewServer()
	pb.RegisterEchoAPIServer(s, api.NewServer())
	log.Println("registered server")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	log.Println("started server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
