package main

import (
	"context"
	"log"
	pb "microservice_grpc/blog/proto"
)

func readBlog(client pb.BlogServiceClient) *pb.Blog {
	log.Println("Reading blog was invoked")

	blog := &pb.BlogId{
		Id: "66384e6b12acbfd8fd976cab",
	}

	res, err := client.ReadBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Failed to read blog: %v\n", err)
	}

	log.Printf("Blog: %v\n", res)

	return res
}
