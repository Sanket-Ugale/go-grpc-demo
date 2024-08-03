package main

import (
	"context"

	pb "github.com/sanket-ugale/grpc-demo/proto"
)

// Ensure the method name matches the RPC method name in your .proto file
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
