package main

import (
	"log"
	pb "microservice_grpc/greet/proto"
)

func (s *Server) GreetManyTimes(request *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", request)

	for i := 0; i < 10; i++ {
		result := "Hello " + request.GetFirstName() + " number " + string(rune(i))
		res := &pb.GreetResponse{
			Result: result,
		}
		stream.Send(res)
	}

	return nil
}
