package main

import (
	"log"
	"time"

	pb "github.com/mudityadev/grpcMessage/proto"
)

// Define the implementation of the SayHelloServerStreaming method
func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	// Log the request names
	log.Printf("Got request with names: %v", req.Names)
	for _, name := range req.Names {
		// Create a response message
		res := &pb.HelloResponse{
			Message: "Good Evening,  " + name,
		}
		// Send the response message to the client
		if err := stream.Send(res); err != nil {
			return err
		}
		// Sleep for 2 seconds to simulate a long-running process
		time.Sleep(2 * time.Second)
	}
	// Return nil to indicate success
	return nil
}
