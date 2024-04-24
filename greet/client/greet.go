package main

import (
	"context"
	"log"
	pb "microservice_grpc/greet/proto"
)

func doGreet(client pb.GreetServiceClient) {
	greet, err := client.Greet(context.Background(), &pb.GreetRequest{FirstName: "John"})
	if err != nil {
		log.Fatalf("Failed to greet: %v\n", err)
	}

	log.Printf("Greet response: %v\n", greet.Result)
}
