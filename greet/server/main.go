package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	pb "microservice_grpc/greet/proto"
	"net"
)

var addr = "localhost:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listen, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Server listening at %v\n", addr)

	var opts []grpc.ServerOption
	tls := true // change that to false if you want

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	server := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(server, &Server{})

	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
