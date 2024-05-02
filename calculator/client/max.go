package main

import (
	"context"
	"io"
	"log"
	pb "microservice_grpc/calculator/proto"
	"time"
)

func doMax(client pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := client.Max(context.Background())

	if err != nil {
		log.Fatalf("Failed to calculate max: %v\n", err)
	}

	waitc := make(chan struct{})
	numbers := []int32{1, 5, 3, 6, 2, 20}

	go func() {
		for _, number := range numbers {
			req := &pb.MaxRequest{
				Number: number,
			}

			if err := stream.Send(req); err != nil {
				log.Fatalf("Failed to send a request: %v\n", err)
			}
			time.Sleep(1 * time.Second)

			log.Printf("Sent: %v\n", number)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Failed to receive response: %v\n", err)
			}

			log.Printf("Max response: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
