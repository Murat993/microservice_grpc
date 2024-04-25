package main

import (
	"context"
	"log"
	pb "microservice_grpc/greet/proto"
	"time"
)

func doLong(client pb.GreetServiceClient) {
	log.Println("Starting to do a LongGreet RPC...")

	requests := []*pb.GreetRequest{
		{
			FirstName: "Stephane",
		},
		{
			FirstName: "Clement",
		},
		{
			FirstName: "Lucas",
		},
	}

	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}

	log.Printf("LongGreet Response: %s\n", res.Result)
}
