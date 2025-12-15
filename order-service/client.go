package main

import (
	"context"
	"fmt"
	"log"
	"time"

	userpb "grpc-microservices/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	client userpb.UserServiceClient
}

func NewUserServiceClient(address string) (*UserServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user service: %v", err)
	}

	client := userpb.NewUserServiceClient(conn)

	return &UserServiceClient{
		client: client,
	}, nil
}

func (c *UserServiceClient) GetUser(userId string) (*userpb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.client.GetUser(ctx, &userpb.GetUserRequest{
		Id: userId,
	})

	if err != nil {
		return nil, err
	}

	return resp.User, nil
}

func (c *UserServiceClient) VerifyUserExists(userId string) bool {
	user, err := c.GetUser(userId)
	if err != nil {
		log.Printf("User verification failed: %v", err)
		return false
	}

	log.Printf("User verified: %s - %s", user.Id, user.Name)
	return true
}
