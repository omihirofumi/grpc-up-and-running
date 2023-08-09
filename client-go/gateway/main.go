package main

import (
	"context"
	gw "ecommerce/gateway/proto/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

var (
	grpcServerEndPoint = "localhost:50051"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	err := gw.RegisterProductInfoHandlerFromEndpoint(ctx, mux,
		grpcServerEndPoint, opts)
	if err != nil {
		log.Fatalf("fail to register gRPC gateway service endpoint: %v", err)
	}

	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("could not setup HTTP endpoint: %v", err)
	}
}
