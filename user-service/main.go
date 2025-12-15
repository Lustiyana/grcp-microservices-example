package main

import (
	"fmt"
	"log"
	"net"

	pb "grpc-microservices/proto/user"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userServer := NewUserServer()

	pb.RegisterUserServiceServer(grpcServer, userServer)

	fmt.Printf("User Service is running on port %s\n", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
