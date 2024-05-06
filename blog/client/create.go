package main

import (
	"context"
	"log"
	pb "microservice_grpc/blog/proto"
)

func createBlog(client pb.BlogServiceClient) string {
	log.Println("Creating blog was invoked")

	blog := &pb.Blog{
		AuthorId: "1",
		Title:    "My first blog",
		Content:  "Content of my first blog",
	}

	res, err := client.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Failed to create blog: %v\n", err)
	}

	log.Printf("Blog created with ID: %v\n", res.Id)

	return res.Id
}
