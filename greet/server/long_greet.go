package main

import (
	"fmt"
	"io"
	"log"
	pb "microservice_grpc/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked with a streaming request")

	result := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		result += fmt.Sprintf("Hello %s!\n", req.GetFirstName())
	}
}
