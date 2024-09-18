package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/viccabezos/go+grpc-YT/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {

	log.Printf("Biderectional streaming started")

	stream, err := client.SayHelloBidirectionalStreaming(context.Background())

	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("could while streaming: %v", err)
			}

			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{Name: name}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending the request: %v", err)
		}
		log.Printf("sent request with name: %v", name)
		time.Sleep(2 * time.Second)

	}
	stream.CloseSend()
	<-waitc
	log.Printf("client streaming finished")

}
