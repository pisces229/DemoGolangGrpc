package main

import (
	"context"
	"demo.golang.grpc.server/grpcWeb/pb"
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"net/http"
)

type runnerServer struct {
}

func (s *runnerServer) Run(ctx context.Context, request *pb.RunnerRequest) (*pb.RunnerResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Println(md)
		if value, ok := md["key"]; ok {
			log.Println(value)
		}
	}
	return &pb.RunnerResponse{Message: "Run:[" + request.Name + "]"}, nil
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
	pb.RegisterRunnerServer(grpcServer, &runnerServer{})
	// Serve gRPC server
	log.Println("Serving gRPC")
	go func() {
		log.Fatalf("failed to serve: %v", grpcServer.Serve(listen))
	}()

	// gRPC web code
	grpcWebServer := grpcweb.WrapServer(
		grpcServer,
		// Enable CORS
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	webServer := &http.Server{
		Handler: grpcWebServer,
		Addr:    fmt.Sprintf("localhost:%d", 8090),
	}
	log.Println("Serving gRPC-Web")
	//log.Fatalln(gatewayServer.ListenAndServe())
	log.Fatalln(webServer.ListenAndServeTLS(certFile, keyFile))
}
