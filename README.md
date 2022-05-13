# PubSubSubscriber

## Subscriber Go Service

This is a service, that demonstrates ProtoBuf using gRPC. 
It is a very basic service, but moving forward will adapt Dapr and the use of PubSub technology to be able to
send messages between this and the PubSubPublisher service.

I have a `.proto` file in this directory, and generate protobuf files using `make proto` via the Makefile

First have `PubSubSubscriber` running

Start both demos with `go run .`
