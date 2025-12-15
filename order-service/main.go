package main

import (
	"fmt"
	"log"
	"net"

	pb "grpc-microservices/proto/order"

	"google.golang.org/grpc"
)

const (
	port            = ":50052"
	userServiceAddr = "localhost:50051"
)

func main() {
	// Initialize User Service Client (untuk komunikasi antar service)
	userClient, err := NewUserServiceClient(userServiceAddr)
	if err != nil {
		log.Printf("Warning: Failed to connect to User Service: %v", err)
		log.Println("Order Service will continue but user verification will be disabled")
	} else {
		log.Println("Successfully connected to User Service")
	}

	// Setup gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderServer := NewOrderServer()

	// Simpan user client untuk digunakan nanti jika diperlukan
	_ = userClient

	pb.RegisterOrderServiceServer(grpcServer, orderServer)

	fmt.Printf("Order Service is running on port %s\n", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
