package main

import (
	"context"
	"demo.golang.grpc.server/grpcWeb/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

// "cert.pem",
// "c:/workspace/Mkcert/localhost+2.pem"
var certFile = "c:/workspace/Mkcert/localhost+2.pem"

func main() {
	doRunner()
}

func doRunner() {
	// Create tls based credential.
	transportCredentials, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(
		"localhost:8080",
		//grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTransportCredentials(transportCredentials),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewRunnerClient(conn)

	md := metadata.Pairs("key", "value")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	response, err := client.Run(ctx, &pb.RunnerRequest{Name: "Golang"})
	if err != nil {
		log.Fatalf("Run fail: %v", err)
	}
	log.Printf("Run: %s", response.GetMessage())
}
