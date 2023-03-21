package main

import (
	"context"

	pb "github.com/mudityadev/grpcMessage/proto"
)

// Define the implementation of the SayHello method
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	// Create a HelloResponse message
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
