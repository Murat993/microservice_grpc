package main

import (
	"context"
	"io"
	"log"
	pb "microservice_grpc/calculator/proto"
)

func doPrimes(client pb.CalculatorServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.PrimeRequest{
		Number: 120,
	}

	stream, err := client.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to calculator many times: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to receive response: %v\n", err)
		}

		log.Printf("Primes response: %v\n", res.Result)
	}

	stream.CloseSend()
}
