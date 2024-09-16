package main

import (
	"context"

	pb "github.com/viccabezos/go+grpc-YT/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{Message: "Hello "}, nil
}
