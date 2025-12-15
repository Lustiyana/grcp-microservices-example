package main

import (
	"context"
	"fmt"
	"log"
	"time"

	orderpb "grpc-microservices/proto/order"
	userpb "grpc-microservices/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to User Service
	userConn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer userConn.Close()

	userClient := userpb.NewUserServiceClient(userConn)

	// Connect to Order Service
	orderConn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}
	defer orderConn.Close()

	orderClient := orderpb.NewOrderServiceClient(orderConn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Test 1: Create User
	fmt.Println("=== Test 1: Create User ===")
	userResp, err := userClient.CreateUser(ctx, &userpb.CreateUserRequest{
		Name:  "John Doe",
		Email: "john@example.com",
		Phone: "08123456789",
	})
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	fmt.Printf("User created: %+v\n\n", userResp.User)

	userId := userResp.User.Id

	// Test 2: Get User
	fmt.Println("=== Test 2: Get User ===")
	getUserResp, err := userClient.GetUser(ctx, &userpb.GetUserRequest{
		Id: userId,
	})
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}
	fmt.Printf("User retrieved: %+v\n\n", getUserResp.User)

	// Test 3: Create Order
	fmt.Println("=== Test 3: Create Order ===")
	orderResp, err := orderClient.CreateOrder(ctx, &orderpb.CreateOrderRequest{
		UserId:      userId,
		ProductName: "Laptop",
		Quantity:    1,
		TotalPrice:  15000000,
	})
	if err != nil {
		log.Fatalf("Failed to create order: %v", err)
	}
	fmt.Printf("Order created: %+v\n\n", orderResp.Order)

	orderId := orderResp.Order.Id

	// Test 4: Create another order
	fmt.Println("=== Test 4: Create Another Order ===")
	orderResp2, err := orderClient.CreateOrder(ctx, &orderpb.CreateOrderRequest{
		UserId:      userId,
		ProductName: "Mouse",
		Quantity:    2,
		TotalPrice:  500000,
	})
	if err != nil {
		log.Fatalf("Failed to create order: %v", err)
	}
	fmt.Printf("Order created: %+v\n\n", orderResp2.Order)

	// Test 5: Get Order
	fmt.Println("=== Test 5: Get Order ===")
	getOrderResp, err := orderClient.GetOrder(ctx, &orderpb.GetOrderRequest{
		Id: orderId,
	})
	if err != nil {
		log.Fatalf("Failed to get order: %v", err)
	}
	fmt.Printf("Order retrieved: %+v\n\n", getOrderResp.Order)

	// Test 6: Get Orders by User ID
	fmt.Println("=== Test 6: Get Orders by User ID ===")
	userOrdersResp, err := orderClient.GetOrdersByUserId(ctx, &orderpb.GetOrdersByUserIdRequest{
		UserId: userId,
	})
	if err != nil {
		log.Fatalf("Failed to get user orders: %v", err)
	}
	fmt.Printf("Total orders for user: %d\n", userOrdersResp.Total)
	for i, order := range userOrdersResp.Orders {
		fmt.Printf("Order %d: %+v\n", i+1, order)
	}
	fmt.Println()

	// Test 7: List All Users
	fmt.Println("=== Test 7: List All Users ===")
	listUsersResp, err := userClient.ListUsers(ctx, &userpb.ListUsersRequest{
		Page:  1,
		Limit: 10,
	})
	if err != nil {
		log.Fatalf("Failed to list users: %v", err)
	}
	fmt.Printf("Total users: %d\n", listUsersResp.Total)
	for i, user := range listUsersResp.Users {
		fmt.Printf("User %d: %+v\n", i+1, user)
	}
	fmt.Println()

	// Test 8: List All Orders
	fmt.Println("=== Test 8: List All Orders ===")
	listOrdersResp, err := orderClient.ListOrders(ctx, &orderpb.ListOrdersRequest{
		Page:  1,
		Limit: 10,
	})
	if err != nil {
		log.Fatalf("Failed to list orders: %v", err)
	}
	fmt.Printf("Total orders: %d\n", listOrdersResp.Total)
	for i, order := range listOrdersResp.Orders {
		fmt.Printf("Order %d: %+v\n", i+1, order)
	}

	fmt.Println("\n=== All tests completed successfully! ===")
}
