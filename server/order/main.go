package main

import (
	"crypto/tls"
	pb "ecommerce/server/order/proto/v1"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

var (
	orderMap = make(map[string]*pb.Order)

	certFile = "./config/server-cert.pem"
	keyFile  = "./config/server-key.pem"
)

func init() {
	orderMap["102"] = &pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00}
	orderMap["103"] = &pb.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00}
	orderMap["104"] = &pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 400.00}
	orderMap["105"] = &pb.Order{Id: "105", Items: []string{"Amazon Echo"}, Destination: "San Jose, CA", Price: 30.00}
	orderMap["106"] = &pb.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"}, Destination: "Mountain View, CA", Price: 300.00}
}

func main() {
	// get port from args
	var port uint
	flag.UintVar(&port, "port", 50052, "Server port to listen on")
	flag.Parse()

	s := &server{orderMap: orderMap}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load key pair: %v", err)
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.NoClientCert,
	}
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewTLS(config)),
		grpc.UnaryInterceptor(orderUnaryServerInterceptor),
		grpc.StreamInterceptor(orderServerStreamInterceptor),
	}

	gs := grpc.NewServer(opts...)
	pb.RegisterOrderManagementServer(gs, s)

	log.Fatal(gs.Serve(lis))

}
