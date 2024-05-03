package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "microservice_grpc/greet/proto"
)

var addr = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to dial: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	//doGreet(client)
	//doGreetManyTimes(client)
	//doLong(client)
	//doGreetEveryone(client)
	doGreetWithDeadline(client)
}
