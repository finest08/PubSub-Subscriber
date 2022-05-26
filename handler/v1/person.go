package handler

import (
	"context"
	"fmt"

	dapr "github.com/dapr/go-sdk/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/finest08/PubSubSubscriber/gen/proto/go/proto/person/v1"
)

// server is used to implement v1.PersonServer.
type PersonServer struct {
	Dapr  dapr.Client
	pb.UnimplementedPersonServiceServer
}

func (p PersonServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	person := req.Person
	
	// publish event
	if err := p.Dapr.PublishEvent(
		context.Background(),
		"pubsub-publish", "restopic", person,
		dapr.PublishEventWithContentType("application/json"),
	); err != nil {
		return &pb.CreateResponse{}, status.Errorf(codes.Aborted, "%s", "error publishing event")
	}
	fmt.Println("Published event: ", person.FirstName)
	return &pb.CreateResponse{Message: "Submission for " + person.FirstName + " " + person.LastName + " posted successfully."}, nil
}
