package main

import (
	"chatservice/api"
	pb "chatservice/proto"
	pgsqlservice "chatservice/service/pgsql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	portAuth = 50051
	portChat = 50052
)

func main() {
	err := pgsqlservice.ConnectToPostgreSQL()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	startServers()
}

func startServers() {
	startAuthService()
}

func startAuthService() {
	fmt.Println("Starting Auth Service")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", portAuth))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go startChatService()
	s := grpc.NewServer()
	server := api.NewAuthServiceServer()
	pb.RegisterAuthServiceServer(s, server)
	s.Serve(lis)
}

func startChatService() {
	fmt.Println("Starting Chat service")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", portChat))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := api.NewChatServiceServer()
	pb.RegisterChatServiceServer(s, server)
	s.Serve(lis)
}
