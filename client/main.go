package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpcexample.com/m/proto"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		log.Fatalf("Failed to start the client %v", err)
	}
	client := pb.NewGreetServiceClient(conn)

	//CallSayHello(client)

	names := &pb.NamesList{
		Names: []string{"Alice", "BoB", "Shristy"},
	}
	CallSayHelloServerStream(client, names)
}
