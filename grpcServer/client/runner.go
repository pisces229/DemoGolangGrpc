package main

import (
	"context"
	"demo.golang.grpc.server/grpcServer/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"time"
)

func doRunner() {
	//doRunnerRun()
	//doRunnerServerStreaming()
	//doRunnerClientStreaming()
	doRunnerBidirectionalStreaming()
}

func doRunnerRun() {
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
	response, err := client.Run(ctx, &pb.RunnerRequest{Message: "Golang"})
	if err != nil {
		log.Fatalf("Run fail: %v", err)
	}
	log.Println(response.GetMessage())
}

func doRunnerServerStreaming() {
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	responseStream, err := client.ServerStreaming(ctx, &pb.RunnerRequest{Message: "Golang"})
	if err != nil {
		log.Fatalf("ServerStreaming fail: %v", err)
	}
	for {
		response, err := responseStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Recv fail: %v", err)
		}
		fmt.Println(response.GetMessage())
	}
	fmt.Println("Done")
}

func doRunnerClientStreaming() {
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientStream, err := client.ClientStreaming(ctx)
	if err != nil {
		log.Fatalf("ClientStreaming fail: %v", err)
	}

	values := []string{"First", "Second", "Third"}
	for _, value := range values {
		clientStream.Send(&pb.RunnerRequest{Message: value})
		time.Sleep(1 * time.Second)
	}
	response, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("CloseAndRecv fail: %v", err)
	}
	fmt.Println(response)
	fmt.Println("Done")
}

func doRunnerBidirectionalStreaming() {
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.BidirectionalStreaming(ctx)
	if err != nil {
		log.Fatalf("BidirectionalStreaming fail: %v", err)
	}

	doneChan := make(chan string)
	go func() {
		values := []string{"First", "Second", "Third"}
		for _, value := range values {
			stream.Send(&pb.RunnerRequest{Message: value})
			time.Sleep(1 * time.Second)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Println(err)
		}
		doneChan <- "Request Done"
	}()
	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Println(err)
			}
			fmt.Println(response)
		}
		doneChan <- "Response Done"
	}()
	//fmt.Println(<-doneChan)
	//fmt.Println(<-doneChan)
	fmt.Println("Done")
}
