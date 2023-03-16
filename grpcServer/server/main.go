package main

import (
	"context"
	runnerPb "demo.golang.grpc.server/grpcServer/pb"
	groupPb "demo.golang.grpc.server/grpcServer/pb/group"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

type runnerServer struct {
}

func (s *runnerServer) Run(ctx context.Context, request *runnerPb.RunnerRequest) (*runnerPb.RunnerResponse, error) {
	return &runnerPb.RunnerResponse{Message: "Run:[" + request.Name + "]"}, nil
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
