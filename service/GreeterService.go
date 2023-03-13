package service

import (
	"context"
	"demo.golang.grpc.server/pb"
)

type GreeterService struct {
}

func (m GreeterService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Server say hello to " + in.GetName()}, nil
}
