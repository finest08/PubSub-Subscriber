package handler

import (
	"context"
	"log"

	dapr "github.com/dapr/go-sdk/client"

	pb "github.com/finest08/PubSubSubscriber/gen/proto/go/proto/person/v1"

)

// server is used to implement v1.PersonServer.
type PersonServer struct {
	Dapr  dapr.Client
	pb.UnimplementedPersonServiceServer
}

// Person implements
func (s PersonServer) Person(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Printf("Received message: \nName: %v %v,\nEmail: %v,\n", in.Person.FirstName, in.Person.LastName, in.Person.Email)
	return &pb.CreateResponse{}, nil
}
