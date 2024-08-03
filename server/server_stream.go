package main

import (
	"log"
	"time"

	pb "github.com/sanket-ugale/grpc-demo/proto"
)

// SayHelloServerStreaming implements the server-side logic for server streaming
func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			log.Printf("error sending response: %v", err)
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
