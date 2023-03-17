package main

import (
	"context"
	"demo.golang.grpc.server/grpcGateway/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// "cert.pem",
// "c:/workspace/Mkcert/localhost+2.pem"
var certFile = "c:/workspace/Mkcert/localhost+2.pem"

func main() {
	doGrpcRunGet()
	doGrpcRunPost()
	doHttpRunGet()
	doHttpRunPost()
}

func doGrpcRunGet() {
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
	response, err := client.RunGet(ctx, &pb.RunnerRequest{Name: "Golang"})
	if err != nil {
		log.Fatalf("RunGet fail: %v", err)
	}
	log.Printf("RunGet: %s", response.GetMessage())
}

func doGrpcRunPost() {
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.RunPost(ctx, &pb.RunnerRequest{Name: "Golang"})
	if err != nil {
		log.Fatalf("RunPost fail: %v", err)
	}
	log.Printf("RunPost: %s", response.GetMessage())
}

func doHttpRunGet() {
	//transport := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	//}
	//resource := "http://localhost:8090/api/Runner/RunGet/Golang"
	resource := "https://localhost:8090/api/Runner/RunGet/Golang"
	request, err := http.NewRequest("GET", resource, nil)
	if err != nil {
		log.Fatalf("NewRequest: %v", err)
	}
	// HTTP headers that start with 'Grpc-Metadata-' are mapped to gRPC metadata after removing prefix 'Grpc-Metadata-'.
	request.Header.Add("Grpc-Metadata-Key", "value")
	client := &http.Client{
		//Transport: transport,
	}
	resp, err := client.Do(request)
	//resp, err := client.Get(resource)
	if err != nil {
		log.Fatalf("Do fail: %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll fail: %v", err)
	}
	fmt.Println(string(body))
}

func doHttpRunPost() {
	//transport := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	//}
	//resource := "http://localhost:8090/api/Runner/RunPost"
	resource := "https://localhost:8090/api/Runner/RunPost"
	client := &http.Client{
		//Transport: transport,
	}
	resp, err := client.Post(
		resource,
		"application/json",
		strings.NewReader("{\"name\": \"Golang\"}"),
	)
	if err != nil {
		log.Fatalf("Do fail: %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll fail: %v", err)
	}
	fmt.Println(string(body))
}
