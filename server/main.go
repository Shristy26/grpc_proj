package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpcexample.com/m/proto"
)

type helloServer struct {
	pb.GreetServiceServer
}

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer()
	//we need to register the greet service we created in proto file
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("started server at %v", lis.Addr())
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}
