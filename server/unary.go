package main

import (
	"context"

	pb "grpcexample.com/m/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "hello from uniary server",
	}, nil
}
