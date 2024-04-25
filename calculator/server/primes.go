package main

import (
	"log"
	pb "microservice_grpc/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", in)

	number := in.GetNumber()
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			res := &pb.PrimeResponse{
				Result: divisor,
			}
			stream.Send(res)
			number = number / divisor
		} else {
			divisor = divisor + 1
		}
	}

	return nil
}
