package api

import (
	"context"
	"testing"

	pb "github.com/micnncim/bazel-go-grpc-playground/pkg/service"
)

func Test_server_Echo(t *testing.T) {
	s := new(server)
	msg := "test"
	resp, err := s.Echo(context.Background(), &pb.EchoRequest{
		Message: msg,
	})
	if err != nil {
		t.Fatalf("error occurred: %+v", err)
	}
	if resp.Message != msg {
		t.Fatalf("test failed: got=%s, want=%s", resp.Message, msg)
	}
}
