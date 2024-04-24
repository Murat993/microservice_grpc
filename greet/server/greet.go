package main

import (
	"context"
	"log"
	pb "microservice_grpc/greet/proto"
)

func (s *Server) Greet(context context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received request: %v\n", request)
	return &pb.GreetResponse{Result: "Hello " + request.FirstName}, nil

}
