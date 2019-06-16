package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/micnncim/bazel-go-grpc-playground/pkg/service"
)

const address = "localhost:8080"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewEchoAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.Echo(ctx, &pb.EchoRequest{
		Message: "Hello",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("received response: %s", resp.Message)
}
