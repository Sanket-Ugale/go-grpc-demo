package main

import (
	"io"
	"log"

	pb "github.com/sanket-ugale/grpc-demo/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error { // Fixed function name and service name
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name, // Added a space after "Hello"
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
