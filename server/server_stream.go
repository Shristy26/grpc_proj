package main

import (
	"context"
	"log"
	"time"

	pb "grpcexample.com/m/proto"
)

func (s *helloServer) SayHelloServerStream(ctx context.Context, req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got request from Names %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}

	return nil

}
