package main

import (
	"context"
	"log"
	"time"

	pb "grpcexample.com/m/proto"
)

func CallSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("[Client Uniary] Could not execute sayHello %v", err)
	}
	log.Printf("[Client Uniary Response]%s", res.Message)
}
