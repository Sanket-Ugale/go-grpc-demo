package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/sanket-ugale/grpc-demo/proto"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional Streaming Started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background()) // Fixed the initialization of the stream
	if err != nil {
		log.Fatalf("Could not start bidirectional streaming: %v", err)
	}
	waitc := make(chan struct{})

	go func() {
		defer close(waitc)
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
			}
			log.Printf("Received message: %s", message.Message) // Print the received message
		}
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}
		log.Printf("Sent request with name: %s", name) // Log the sent request
		time.Sleep(2 * time.Second)
	}
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error while closing send: %v", err)
	}
	<-waitc
	log.Printf("Bidirectional streaming finished")
}
