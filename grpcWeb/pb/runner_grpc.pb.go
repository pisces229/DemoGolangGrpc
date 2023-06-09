// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.1
// source: runner.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Runner_Run_FullMethodName = "/runner.Runner/Run"
)

// RunnerClient is the client API for Runner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RunnerClient interface {
	Run(ctx context.Context, in *RunnerRequest, opts ...grpc.CallOption) (*RunnerResponse, error)
}

type runnerClient struct {
	cc grpc.ClientConnInterface
}

func NewRunnerClient(cc grpc.ClientConnInterface) RunnerClient {
	return &runnerClient{cc}
}

func (c *runnerClient) Run(ctx context.Context, in *RunnerRequest, opts ...grpc.CallOption) (*RunnerResponse, error) {
	out := new(RunnerResponse)
	err := c.cc.Invoke(ctx, Runner_Run_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunnerServer is the server API for Runner service.
// All implementations should embed UnimplementedRunnerServer
// for forward compatibility
type RunnerServer interface {
	Run(context.Context, *RunnerRequest) (*RunnerResponse, error)
}

// UnimplementedRunnerServer should be embedded to have forward compatible implementations.
type UnimplementedRunnerServer struct {
}

func (UnimplementedRunnerServer) Run(context.Context, *RunnerRequest) (*RunnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}

// UnsafeRunnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RunnerServer will
// result in compilation errors.
type UnsafeRunnerServer interface {
	mustEmbedUnimplementedRunnerServer()
}

func RegisterRunnerServer(s grpc.ServiceRegistrar, srv RunnerServer) {
	s.RegisterService(&Runner_ServiceDesc, srv)
}

func _Runner_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Runner_Run_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Run(ctx, req.(*RunnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Runner_ServiceDesc is the grpc.ServiceDesc for Runner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Runner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "runner.Runner",
	HandlerType: (*RunnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _Runner_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runner.proto",
}
