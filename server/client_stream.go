package main

import (
	"io"
	"log"

	pb "github.com/mudityadev/grpcMessage/proto"
)

// Define the implementation of the SayHelloClientStreaming method
func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	// Initialize an empty slice to hold messages
	var messages []string

	// Loop through incoming messages until the stream is closed
	for {
		// Receive a message from the stream
		req, err := stream.Recv()
		if err == io.EOF {
			// If the end of the stream is reached, send a response with the collected messages
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}

		// Log the received message
		log.Printf("Got request  name: %v", req.Name)

		// Add a new message to the slice
		messages = append(messages, "Hello "+req.Name)
	}

	// Return nil to indicate success
	return nil
}
