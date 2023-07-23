package main

import (
	"context"
	"log"

	pb "github.com/mzzz-zzm/experimental/blazorapp1/go/proto/greet"
)

func (s *Server) Greet(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.HelloResponse{
		Result: "HelloWorld! from server for " + in.FirstName,
	}, nil
}
