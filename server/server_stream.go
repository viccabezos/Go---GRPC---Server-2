package main

import (
	"log"
	"time"

	pb "github.com/viccabezos/go+grpc-YT/proto"
)

// server receive a request with a list of names and return a stream of HelloResponse
func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {

	log.Printf("got request with names: %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{Message: "Hello " + name}
		if err := stream.Send(res); err != nil {
			return err
		}
		//add delay to simulate a long running operation
		time.Sleep(2 * time.Second)
	}

	return nil
}
