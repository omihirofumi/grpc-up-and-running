package main

import (
	"crypto/tls"
	"crypto/x509"
	pb "ecommerce/server/order/proto/v1"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
)

var (
	orderMap = make(map[string]*pb.Order)

	certFile = "./config/server-cert.pem"
	keyFile  = "./config/server-key.pem"
	caFile   = "../config/ca-cert.pem"

	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid credentials")
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
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(caFile)
	if err != nil {
		log.Fatalf("could not read ca certificate: %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append ca certificate: %v", err)
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequestClientCert,
		ClientCAs:    certPool,
	}
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewTLS(config)),
		grpc.ChainUnaryInterceptor(
			ensureValidBasicCredentials,
			orderUnaryServerInterceptor,
		),
		grpc.StreamInterceptor(orderServerStreamInterceptor),
	}

	gs := grpc.NewServer(opts...)
	pb.RegisterOrderManagementServer(gs, s)

	log.Fatal(gs.Serve(lis))

}
