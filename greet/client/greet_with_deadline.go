package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	pb "microservice_grpc/greet/proto"
	"time"
)

func doGreetWithDeadline(client pb.GreetServiceClient) {
	log.Println("Starting to do a GreetWithDeadline RPC...")
	req := &pb.GreetRequest{
		FirstName: "Stephane",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.GreetWithDeadline(ctx, req)

	if err != nil {
		if status.Code(err) == codes.DeadlineExceeded {
			log.Fatalln("Timeout was hit! Deadline was exceeded")
		}
		log.Fatalf("Error while calling GreetWithDeadline RPC: %v", err)
	}

	log.Printf("Response from GreetWithDeadline: %v", res.Result)
}
