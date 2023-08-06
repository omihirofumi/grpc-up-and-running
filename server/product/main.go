package main

import (
	pb "ecommerce/server/product/proto/v1"
	"fmt"
	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/zpages"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

const port = 50051

func main() {
	go func() {
		mux := http.NewServeMux()
		zpages.Handle(mux, "/debug")
		log.Fatal(http.ListenAndServe("127.0.0.1:8001", mux))
	}()

	view.RegisterExporter(&exporter.PrintExporter{})

	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))

	pb.RegisterProductInfoServer(s, &server{})

	log.Printf("Starting gRPC listener on port %d", port)

	log.Fatal(s.Serve(lis))
}
