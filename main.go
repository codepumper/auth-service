package main

import (
	"log"
	"net"

	"auth-service/handlers"
	"auth-service/proto"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	server := grpc.NewServer()
	authService := handlers.NewAuthService()
	auth.RegisterAuthServiceServer(server, authService)

	log.Println("Auth Service is running on port 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
