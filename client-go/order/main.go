package main

import (
	"context"
	pb "ecommerce/client/order/proto/v1"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"log"
	"time"
)

const (
	address   = "localhost:50052"
	localhost = "localhost"
)

var certFile = "../server/server-cert.pem"

func main() {
	creds, err := credentials.NewClientTLSFromFile(certFile, localhost)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrderManagementClient(conn)

	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"kn", "vn",
	)
	mdCtx := metadata.NewOutgoingContext(context.Background(), md)

	ctxA := metadata.AppendToOutgoingContext(mdCtx, "k1", "v1", "k2", "v2")

	var header, trailer metadata.MD

	order1 := pb.Order{Id: "101", Items: []string{"A", "B"}, Destination: "Japan", Price: 100}
	res, err := client.AddOrder(ctxA, &order1, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatal(err)
	}

	log.Print("AddOrder Response -> ", res.Value)

	if t, ok := header["timestamp"]; ok {
		log.Printf("timestamp from header:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("timestamp expected but doesn't exist in header")
	}
	if l, ok := trailer["location"]; ok {
		log.Printf("location from trailer:\n")
		for i, e := range l {
			fmt.Printf(" %d . %s\n", i, e)
		}
	} else {
		log.Fatal("location expected but doesn't exist in trailer")
	}

	searchStream, _ := client.SearchOrders(ctxA, &wrapperspb.StringValue{Value: "Google"})
	for {
		searchOrder, err := searchStream.Recv()
		if err == io.EOF {
			log.Print("EOF")
			break
		}

		if err == nil {
			log.Print("Search Result : ", searchOrder)
		}
	}

	// Update Orders
	updOrder1 := pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Google Pixel Book"}, Destination: "Mountain View, CA", Price: 1100.00}
	updOrder2 := pb.Order{Id: "103", Items: []string{"Apple Watch S4", "Mac Book Pro", "iPad Pro"}, Destination: "San Jose, CA", Price: 2800.00}
	updOrder3 := pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub", "iPad Mini"}, Destination: "Mountain View, CA", Price: 2200.00}

	updateStream, _ := client.UpdateOrders(mdCtx)

	_ = updateStream.Send(&updOrder1)
	_ = updateStream.Send(&updOrder2)
	_ = updateStream.Send(&updOrder3)

	updateRes, _ := updateStream.CloseAndRecv()
	log.Printf("Update Orders Res : ", updateRes)
}
