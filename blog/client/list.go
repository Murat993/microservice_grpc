package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"io"
	"log"
	pb "microservice_grpc/blog/proto"
)

func listBlog(client pb.BlogServiceClient) {
	log.Println("List blog was invoked")

	stream, err := client.ListBlog(context.Background(), &empty.Empty{})

	if err != nil {
		log.Fatalf("Failed to list blog: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive data: %v\n", err)
		}
		log.Printf("Blog: %v\n", res)
	}
}
