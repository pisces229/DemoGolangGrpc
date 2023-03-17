package main

import (
	"context"
	"demo.golang.grpc.server/grpcGateway/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"net/http"
)

type server struct {
	pb.UnimplementedRunnerServer
}

func (s *server) RunGet(ctx context.Context, request *pb.RunnerRequest) (*pb.RunnerResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Println(md)
		if value, ok := md["key"]; ok {
			log.Println(value)
		}
	}
	return &pb.RunnerResponse{Message: "RunGet:[" + request.Name + "]"}, nil
}

func (s *server) RunPost(ctx context.Context, request *pb.RunnerRequest) (*pb.RunnerResponse, error) {
	return &pb.RunnerResponse{Message: "RunPost:[" + request.Name + "]"}, nil
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
	pb.RegisterRunnerServer(grpcServer, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC")
	go func() {
		log.Fatalln(grpcServer.Serve(listen))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:8080",
		//grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTransportCredentials(transportCredentials),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	serveMux := runtime.NewServeMux()
	err = pb.RegisterRunnerHandler(context.Background(), serveMux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	gatewayServer := &http.Server{
		Addr:    ":8090",
		Handler: serveMux,
	}
	log.Println("Serving gRPC-Gateway")
	//log.Fatalln(gatewayServer.ListenAndServe())
	log.Fatalln(gatewayServer.ListenAndServeTLS(certFile, keyFile))
}
