package main

import (
	"demo.golang.grpc.server/pb"
	"demo.golang.grpc.server/service"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func main() {
	fmt.Println("server...")
	// Create gRPC Server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create tls based credential.
	creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")

	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	log.Println("gRPC server is running.")
	pb.RegisterGreeterServer(grpcServer, service.GreeterService{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("...server")
}
