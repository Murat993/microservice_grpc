package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	pb "microservice_grpc/greet/proto"
	"time"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("GreetWithDeadline function was invoked with a deadline")

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("The client cancelled the request: %v", ctx.Err())
			return nil, status.Error(codes.Canceled, "The client cancelled the request")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{Result: "Hello " + in.FirstName}, nil
}
