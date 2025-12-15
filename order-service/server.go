package main

import (
	"context"
	"fmt"
	"sync"

	pb "grpc-microservices/proto/order"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	orders map[string]*pb.Order
	mu     sync.RWMutex
}

func NewOrderServer() *OrderServer {
	return &OrderServer{
		orders: make(map[string]*pb.Order),
	}
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Validate input
	if req.UserId == "" || req.ProductName == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id and product_name are required")
	}

	if req.Quantity <= 0 {
		return nil, status.Error(codes.InvalidArgument, "quantity must be greater than 0")
	}

	// Create order
	order := &pb.Order{
		Id:          uuid.New().String(),
		UserId:      req.UserId,
		ProductName: req.ProductName,
		Quantity:    req.Quantity,
		TotalPrice:  req.TotalPrice,
		Status:      "pending",
	}

	s.orders[order.Id] = order

	fmt.Printf("Order created: %s - %s (User: %s)\n", order.Id, order.ProductName, order.UserId)

	return &pb.OrderResponse{
		Order:   order,
		Message: "Order created successfully",
	}, nil
}

func (s *OrderServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	order, exists := s.orders[req.Id]
	if !exists {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	return &pb.OrderResponse{
		Order:   order,
		Message: "Order retrieved successfully",
	}, nil
}

func (s *OrderServer) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var orders []*pb.Order
	for _, order := range s.orders {
		orders = append(orders, order)
	}

	return &pb.ListOrdersResponse{
		Orders: orders,
		Total:  int32(len(orders)),
	}, nil
}

func (s *OrderServer) GetOrdersByUserId(ctx context.Context, req *pb.GetOrdersByUserIdRequest) (*pb.ListOrdersResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var orders []*pb.Order
	for _, order := range s.orders {
		if order.UserId == req.UserId {
			orders = append(orders, order)
		}
	}

	return &pb.ListOrdersResponse{
		Orders: orders,
		Total:  int32(len(orders)),
	}, nil
}
