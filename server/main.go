package main

import (
	"log"
	"net"
	pb "github.com/mudityadev/grpcMessage/proto"
	"google.golang.org/grpc"
)

// Define the port to listen on
const (
	port = ":8080"
)

// Define a struct to implement the GreetServiceServer interface
type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	// Listen on the port specified in the const
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the greet service with the server
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	// Start the server and log the address
	log.Printf("Server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
