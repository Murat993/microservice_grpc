package main

import (
	"context"
	"log"
	pb "microservice_grpc/calculator/proto"
)

func (s *Server) Sum(context context.Context, request *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Received request: %v\n", request)
	return &pb.SumResponse{Result: request.FirstNumber + request.SecondNumber}, nil
}
