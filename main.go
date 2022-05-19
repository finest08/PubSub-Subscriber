// Package main implements a server for Greeter service.
package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/finest08/PubSubSubscriber/gen/proto/go/proto/person/v1"
	"github.com/finest08/PubSubSubscriber/handler/v1"
)

func main() {
	s := grpc.NewServer()
	defer s.Stop()
	
	pb.RegisterPersonServiceServer(s, handler.PersonServer{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}



