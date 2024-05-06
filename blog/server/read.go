package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	pb "microservice_grpc/blog/proto"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Println("ReadBlog")

	oid, err := primitive.ObjectIDFromHex(in.GetId())

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid ObjectID")
	}

	data := &BlogItem{}

	filter := primitive.D{{Key: "_id", Value: oid}}

	err = collection.FindOne(ctx, filter).Decode(data)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Blog not found")
	}

	return &pb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}, nil
}
