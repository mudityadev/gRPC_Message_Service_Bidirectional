package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/mudityadev/grpcMessage/proto"
)

// callSayHelloBidirectionalStream is a function that sends a list of names to the server
// using the SayHelloBidirectionalStreaming method of the GreetServiceClient interface.
// It then receives a stream of messages from the server and logs them to the console.
func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	// Log the start of the bidirectional streaming process
	log.Printf("Bidirectional Streaming started")

	// Call the SayHelloBidirectionalStreaming method on the client
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	// Create a channel to wait for the stream to complete
	waitc := make(chan struct{})

	// Create a goroutine to receive messages from the server
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			// Log each message received from the server
			log.Println(message)
		}
		// Close the wait channel when the stream is complete
		close(waitc)
	}()

	// Send each name in the NamesList to the server using the Send method of the stream
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		// Add a delay to simulate a long-running process
		time.Sleep(2 * time.Second)
	}

	// Close the send stream
	stream.CloseSend()

	// Wait for the stream to complete
	<-waitc

	// Log the end of the bidirectional streaming process
	log.Printf("Bidirectional Streaming finished")
}
 