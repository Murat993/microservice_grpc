package main

import (
	"io"
	"log"
	pb "microservice_grpc/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked with a streaming request")

	sum := 0.0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: sum / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		sum += float64(req.GetNumber())
		count++
	}

	return nil
}
