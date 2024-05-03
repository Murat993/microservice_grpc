package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	pb "microservice_grpc/calculator/proto"
)

func doSqrt(client pb.CalculatorServiceClient) {
	log.Println("Starting to do a Unary Square Root RPC...")

	res, err := client.Sqrt(context.Background(), &pb.SqrtRequest{Number: -2})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %v\n", e.Message())

			log.Printf("Error code from server: %v\n", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				return
			}

		} else {
			log.Fatalf("A non gRPC error occurred: %v\n", err)
		}
	}

	log.Printf("Response from Square Root: %v", res.Result)
}
