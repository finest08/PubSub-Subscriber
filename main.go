// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/finest08/PubSubSubscriber/gen/proto/go/proto/person/v1"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement v1.PersonServer.
type server struct {
	pb.UnimplementedPersonServiceServer
}

// Person implements
func (s *server) Person(ctx context.Context, in *pb.PersonRequest) (*pb.PersonResponse, error) {
	log.Printf("Received message: \nName: %v %v,\nEmail: %v,\nOccupation: %v,\nAge: %v", in.GetFirstName(), in.GetLastName(), in.GetEmail(), in.GetOccupation(), in.GetAge())
	return &pb.PersonResponse{Message: "Received Person Details: " + in.GetFirstName() + " " + in.GetLastName()}, nil
}

func main() {

	lis, err := net.Listen("tcp",fmt.Sprintf(":%d", *port) )
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPersonServiceServer(s, &server{})
	log.Printf("Subscribing service listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
