package main

import (
	"io"
	"log"
	pb "microservice_grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked with a streaming request")
	var max int32 = 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}
		if req.GetNumber() > max {
			max = req.GetNumber()
			if err := stream.Send(&pb.MaxResponse{Result: max}); err != nil {
				log.Fatalf("Failed to send a response: %v\n", err)
			}
		}
	}
}
