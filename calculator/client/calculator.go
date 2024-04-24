package main

import (
	"context"
	"log"
	pb "microservice_grpc/calculator/proto"
)

func doSum(client pb.CalculatorServiceClient) {
	sum, err := client.Sum(context.Background(), &pb.SumRequest{FirstNumber: 1, SecondNumber: 2})
	if err != nil {
		log.Fatalf("Failed to calculate: %v\n", err)
	}

	log.Printf("Sum response: %d\n", sum.Result)
}
