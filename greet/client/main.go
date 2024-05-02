package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "microservice_grpc/greet/proto"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Failed to dial: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	//doGreet(client)
	//doGreetManyTimes(client)
	//doLong(client)
	doGreetEveryone(client)
}
