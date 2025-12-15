package main

import (
	"context"
	"fmt"
	"sync"

	pb "grpc-microservices/proto/user"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	users map[string]*pb.User
	mu    sync.RWMutex
}

func NewUserServer() *UserServer {
	return &UserServer{
		users: make(map[string]*pb.User),
	}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Validate input
	if req.Name == "" || req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "name and email are required")
	}

	// Create user
	user := &pb.User{
		Id:    uuid.New().String(),
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	s.users[user.Id] = user

	fmt.Printf("User created: %s - %s\n", user.Id, user.Name)

	return &pb.UserResponse{
		User:    user,
		Message: "User created successfully",
	}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[req.Id]
	if !exists {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &pb.UserResponse{
		User:    user,
		Message: "User retrieved successfully",
	}, nil
}

func (s *UserServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var users []*pb.User
	for _, user := range s.users {
		users = append(users, user)
	}

	return &pb.ListUsersResponse{
		Users: users,
		Total: int32(len(users)),
	}, nil
}
