package main

import (
	pb "ecommerce/server/proto/v1"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})

	log.Printf("Starting gRPC listener on port %d", port)

	log.Fatal(s.Serve(lis))
}
