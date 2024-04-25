package main

import (
	"context"
	"log"
	pb "microservice_grpc/calculator/proto"
)

func doAvg(client pb.CalculatorServiceClient) {
	stream, err := client.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Avg RPC: %v", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		req := &pb.AvgRequest{
			Number: number,
		}

		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Error while sending request to Avg RPC: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Avg RPC: %v", err)
	}

	log.Printf("Response from Avg: %v", res.GetResult())
}
