package main

import (
	"context"
	groupPb "demo.golang.grpc.server/grpcServer/pb/group"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

func doGroup() {
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
	client := groupPb.NewGroupClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.Do(ctx, &groupPb.GroupRequest{Name: "Golang"})
	if err != nil {
		log.Fatalf("Do fail: %v", err)
	}
	log.Printf("Do: %s", response.GetMessage())
}
