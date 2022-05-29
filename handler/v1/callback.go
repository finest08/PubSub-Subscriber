package handler

import (
	"context"
	"fmt"

	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"

	pbpers "github.com/finest08/PubSubSubscriber/gen/proto/go/proto/person/v1"
)

type CallbackServer struct {
	PersonServer PersonServer
	pb.UnimplementedAppCallbackServer
}

// Dapr will call this method to get the list of topics the app wants to subscribe to.
func (p CallbackServer) ListTopicSubscriptions(ctx context.Context, in *emptypb.Empty) (*pb.ListTopicSubscriptionsResponse, error) {

	fmt.Println("ListTopicSubscriptions", in)
	// return &pb.ListTopicSubscriptionsResponse{
	// 	Subscriptions: []*pb.TopicSubscription{{
	// 		PubsubName: "pubsub-publish",
	// 		Topic:     "my-topic",
	// 		Routes:     &pb.TopicRoutes{
	// 			Rules: []*pb.TopicRule{
	// 				{
	// 					Match: `event.data.type == "update"`,
	// 					Path:  "/update",
	// 				},
	// 			},
	// 			Default: "/create"},
	// 		}},
	// }, nil


	fmt.Println("ListTopicSubscriptions")
	return &pb.ListTopicSubscriptionsResponse{
		Subscriptions: []*pb.TopicSubscription{{
			PubsubName: "pubsub-publish",
			Topic:      "my-topic",
			Routes:     &pb.TopicRoutes{Default: "/create"},
		}},
	}, nil
}

// OnTopicEvent is fired for events subscribed to.
// Dapr sends published messages in a CloudEvents 0.3 envelope.
func (p CallbackServer) OnTopicEvent(ctx context.Context, in *pb.TopicEventRequest) (*pb.TopicEventResponse, error) {

	fmt.Println("OnTopicEvent:", in.Path, string(in.Data))
	// json event data -> event.EventData
	var per pbpers.Person
	if err := protojson.Unmarshal(in.Data, &per); err != nil {
		return &pb.TopicEventResponse{Status: pb.TopicEventResponse_DROP},
			status.Errorf(codes.Aborted, "issue unmarshalling data: %v", err)
	}

	// fmt.Println(&per)

switch in.Path {
	case "/create":
		fmt.Println("Switch: /create: ", per.FirstName)
		
	case "/update":
		fmt.Println("Switch: /update: ", per.FirstName)
	default:
		return &pb.TopicEventResponse{},
			status.Errorf(codes.Aborted, "unexpected path in OnTopicEvent: %s", in.Path)
	}

	return &pb.TopicEventResponse{}, nil
}