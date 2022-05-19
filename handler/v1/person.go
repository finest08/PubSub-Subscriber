package handler

import (
	"context"
	"log"

	pb "github.com/finest08/PubSubSubscriber/gen/proto/go/proto/person/v1"
)

// server is used to implement v1.PersonServer.
type PersonServer struct {
	pb.UnimplementedPersonServiceServer
}

// Person implements
func (s PersonServer) Person(ctx context.Context, in *pb.PersonRequest) (*pb.PersonResponse, error) {
	log.Printf("Received message: \nName: %v %v,\nEmail: %v,\nOccupation: %v,\nAge: %v", in.GetFirstName(), in.GetLastName(), in.GetEmail(), in.GetOccupation(), in.GetAge())
	return &pb.PersonResponse{Message: "Received Person Details: " + in.GetFirstName() + " " + in.GetLastName()}, nil
}
