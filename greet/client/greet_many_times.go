package main

import (
	"context"
	"io"
	"log"
	pb "microservice_grpc/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Clement",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to greet many times: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to receive response: %v\n", err)
		}

		log.Printf("GreetManyTimes response: %v\n", res.Result)
	}

	stream.CloseSend()

}
