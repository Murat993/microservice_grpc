package main

import (
	"context"
	"io"
	"log"
	pb "microservice_grpc/greet/proto"
	"time"
)

func doGreetEveryone(client pb.GreetServiceClient) {
	log.Println("Starting to do a GreetEveryone RPC...")

	stream, err := client.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while calling GreetEveryone: %v", err)
	}

	requests := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "Stephane"},
		{FirstName: "Lucas"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
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
				log.Printf("Error while receiving: %v", err)
				break
			}

			log.Printf("Received: %v\n", res.GetResult())
		}

		close(waitc)
	}()

	<-waitc
}
