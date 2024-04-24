package main

import (
	"google.golang.org/grpc"
	"log"
	pb "microservice_grpc/calculator/proto"
	"net"
)

var addr = "localhost:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	listen, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Server listening at %v\n", addr)

	server := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(server, &Server{})

	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
