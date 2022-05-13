// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/finest08/PubSubSubscriber/gen/proto/go/proto/greeting/v1"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received message: %v", in.GetName())
	// return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
	return &pb.HelloReply{Message: "Greeting received. LOUD AND CLEAR! "}, nil
}

// Person implements
func (s *server) Person(ctx context.Context, in *pb.PersonRequest) (*pb.PersonReply, error) {
	log.Printf("Received message: \n%v,\n%v,\n%s", in.GetName(), in.GetOccupation(), in.GetAge())
	return &pb.PersonReply{Message: "Received Person Details: " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("Subscribing service listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
