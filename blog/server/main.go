package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	pb "microservice_grpc/blog/proto"
	"net"
)

var collection *mongo.Collection
var addr = "localhost:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatalf("Failed to create mongo client: %v\n", err)
	}

	if err = client.Connect(context.Background()); err != nil {
		log.Fatalf("Failed to connect to mongo: %v\n", err)
	}

	collection = client.Database("mydb").Collection("blog")

	listen, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Server listening at %v\n", addr)

	server := grpc.NewServer()
	pb.RegisterBlogServiceServer(server, &Server{})

	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
