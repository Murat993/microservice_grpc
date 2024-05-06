package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	pb "microservice_grpc/blog/proto"
)

func (s *Server) ListBlog(in *empty.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Println("ListBlog")

	cursor, err := collection.Find(context.Background(), primitive.D{})

	if err != nil {
		return status.Errorf(codes.Internal, "Unknown internal error")
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		data := &BlogItem{}
		err := cursor.Decode(data)
		if err != nil {
			return status.Errorf(codes.Internal, "Error decoding data")
		}
		stream.Send(documentToBlog(data))
	}

	if err = cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, "Unknown internal error")
	}

	return nil
}
