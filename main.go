package main

import (
	"log"
	"net"
	"time"

	daprpb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	dapr "github.com/dapr/go-sdk/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/finest08/PubSubSubscriber/gen/proto/go/proto/person/v1"
	"github.com/finest08/PubSubSubscriber/handler/v1"
)

func main() {
	// initialise Dapr client using DAPR_GRPC_PORT env var
	// N.B. sleep briefly to give the dapr service time to initialise
	time.Sleep(2 * time.Second)
	client, err := dapr.NewClient()
	if err != nil {
		log.Fatalf("failed to initialise Dapr client: %v", err)
	}
	defer client.Close()

	grpcSrv := grpc.NewServer()
	defer grpcSrv.Stop()         // stop server on exit
	reflection.Register(grpcSrv) // for postman
	
	h := &handler.PersonServer{
		Dapr:  client,
	}
	pb.RegisterPersonServiceServer(grpcSrv, h)

	ch := handler.CallbackServer{}
	daprpb.RegisterAppCallbackServer(grpcSrv, ch)
	
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}



