package main

import (
	"context"
	pb "ecommerce/client/product/proto/v1"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {

	args := os.Args
	var hostname, port string
	if len(args) < 3 {
		hostname = "localhost"
		port = "50051"
	} else {
		hostname = args[1]
		port = args[2]
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", hostname, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect grpc server: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductInfoClient(conn)
	id, err := client.AddProduct(context.Background(), &pb.Product{Name: "apple", Price: 100})
	if err != nil {
		log.Println("failed to add product")
	}
	log.Println(id)
}
