package main

import (
	"context"
	runnerPb "demo.golang.grpc.server/grpcServer/pb"
	groupPb "demo.golang.grpc.server/grpcServer/pb/group"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"net"
	"time"
)

type runnerServer struct {
}

func (s *runnerServer) Run(ctx context.Context, request *runnerPb.RunnerRequest) (*runnerPb.RunnerResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Println(md)
		if value, ok := md["key"]; ok {
			log.Println(value)
		}
	}
	return &runnerPb.RunnerResponse{Message: "Run:[" + request.Message + "]"}, nil
}

func (s *runnerServer) ServerStreaming(request *runnerPb.RunnerRequest, responseStream runnerPb.Runner_ServerStreamingServer) error {
	var error error
	log.Println(request.Message)
	values := []string{"First", "Second", "Third"}
	for _, value := range values {
		result := &runnerPb.RunnerResponse{Message: value}
		if err := responseStream.Send(result); err != nil {
			log.Println("failed to Send:", err)
			error = err
			break
		}
		time.Sleep(1 * time.Second)
	}
	return error
}

func (s *runnerServer) ClientStreaming(requestStream runnerPb.Runner_ClientStreamingServer) error {
	var error error
	for {
		request, err := requestStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("failed to Recv:", err)
			error = err
			break
		}
		log.Println(request.Message)
	}
	if err := requestStream.SendAndClose(&runnerPb.RunnerResponse{Message: "ok"}); err != nil {
		log.Println("failed to SendAndClose:", err)
		error = err
	}
	return error
}

func (s *runnerServer) BidirectionalStreaming(stream runnerPb.Runner_BidirectionalStreamingServer) error {
	var error error
	doneChan := make(chan string)
	go func() {
		for {
			request, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Println("failed to Recv:", err)
				error = err
				break
			}
			log.Printf(request.Message)
		}
		doneChan <- "Request Done"
	}()
	go func() {
		values := []string{"First", "Second", "Third"}
		for _, value := range values {
			stream.Send(&runnerPb.RunnerResponse{Message: value})
			time.Sleep(1 * time.Second)
		}
		doneChan <- "Response Done"
	}()
	fmt.Println(<-doneChan)
	fmt.Println(<-doneChan)
	fmt.Println("Done")
	return error
}

type groupServer struct {
}

func (s *groupServer) Do(ctx context.Context, request *groupPb.GroupRequest) (*groupPb.GroupResponse, error) {
	return &groupPb.GroupResponse{Message: "Do:[" + request.Name + "]"}, nil
}

// "cert.pem"
// "c:/workspace/Mkcert/localhost+2.pem"
var certFile = "c:/workspace/Mkcert/localhost+2.pem"

// "key.pem"
// "c:/workspace/Mkcert/localhost+2-key.pem"
var keyFile = "c:/workspace/Mkcert/localhost+2-key.pem"

func main() {
	// Create tls based credential.
	transportCredentials, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}

	// Create a listener on TCP port
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// Create a gRPC server object
	//grpcServer := grpc.NewServer()
	grpcServer := grpc.NewServer(grpc.Creds(transportCredentials))
	runnerPb.RegisterRunnerServer(grpcServer, &runnerServer{})
	groupPb.RegisterGroupServer(grpcServer, &groupServer{})
	// Serve gRPC server
	log.Println("Serving gRPC")
	log.Fatalln(grpcServer.Serve(listen))
}
