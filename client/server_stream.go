package main

import (
	"context"
	"io"
	"log"

	pb "github.com/viccabezos/go+grpc-YT/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {

	log.Printf("streaming has started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	// process the stream
	for {
		//recieve the message
		message, err := stream.Recv()
		if err == io.EOF {
			//check for end of file
			break
		}
		if err != nil {
			//check errors
			log.Fatalf("could while streaming: %v", err)
		}
		log.Println(message)
	}
	log.Printf("streaaming finished")

}
