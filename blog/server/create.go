package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	pb "microservice_grpc/blog/proto"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Println("CreateBlog invoked")

	data := &BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	result, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to insert blog: %v", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(codes.Internal, "Failed to convert OID")
	}

	return &pb.BlogId{Id: oid.Hex()}, nil
}
