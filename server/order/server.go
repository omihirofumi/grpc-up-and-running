package main

import (
	"context"
	pb "ecommerce/server/order/proto/v1"
	"fmt"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"log"
	"strings"
	"sync"
)

const (
	port           = 50052
	orderBatchSize = 3
)

type server struct {
	orderMap map[string]*pb.Order
	mu       sync.RWMutex
	pb.UnimplementedOrderManagementServer
}

func (s *server) AddOrder(ctx context.Context, in *pb.Order) (*wrapperspb.StringValue, error) {
	if in.Id == "-1" {
		log.Printf("Order ID is invalid! -> Received Order ID %s",
			in.Id)

		errorStatus := status.New(codes.InvalidArgument,
			"Invalid information received")
		ds, err := errorStatus.WithDetails(
			&epb.BadRequest_FieldViolation{
				Field:       "ID",
				Description: fmt.Sprintf("Order ID received not valid %s", in.Id),
			})
		if err != nil {
			return nil, errorStatus.Err()
		}
		return nil, ds.Err()
	}

	log.Printf("Order Adding ID: %v", in.Id)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.orderMap[in.Id] = in
	return &wrapperspb.StringValue{Value: "Order Added ID: " + in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) GetOrder(ctx context.Context, in *wrapperspb.StringValue) (*pb.Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	order, exists := s.orderMap[in.Value]
	if exists {
		return order, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Order does not exists. ID: %v", in)
}

func (s *server) SearchOrders(in *wrapperspb.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for key, order := range s.orderMap {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(itemStr, in.Value) {
				err := stream.Send(order)
				if err != nil {
					return status.Errorf(codes.Internal, "error sending message to stream: %v", err)
				}
				log.Print("Matching Order Found: ", key)
			}
		}
	}
	return status.New(codes.OK, "").Err()
}

func (s *server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	orderStr := "Updated Order IDs: "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wrapperspb.StringValue{Value: "Orders Update done: " + orderStr})
		}

		if err != nil {
			return err
		}
		s.mu.Lock()
		s.orderMap[order.Id] = order
		s.mu.Unlock()

		log.Printf("Order ID: %s - updated", order.Id)
		orderStr += order.Id + ", "
	}
}

func (s *server) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {
	bacthMaker := 1
	var combinedShipmentMap = make(map[string]*pb.CombinedShipment)
	for {
		orderId, err := stream.Recv()
		log.Printf("Reading from client: %s", orderId)
		if err == io.EOF {
			log.Printf("Client closed: %s", orderId)
			for _, shipment := range combinedShipmentMap {
				if err := stream.Send(shipment); err != nil {
					return err
				}
			}
		}
		if err != nil {
			log.Print(err)
			return err
		}

		s.mu.RLock()
		ord, exists := s.orderMap[orderId.Value]
		s.mu.RUnlock()

		if !exists {
			log.Printf("Does not exists ID: %s", orderId.Value)
			continue
		}

		dest := ord.Destination
		shipment, exists := combinedShipmentMap[dest]

		if exists {
			shipment.OrdersList = append(shipment.OrdersList, ord)
			combinedShipmentMap[dest] = shipment
		} else {
			combShipment := &pb.CombinedShipment{
				Id:     fmt.Sprintf("cmb - %s", dest),
				Status: "Processed!",
			}
			combShipment.OrdersList = []*pb.Order{ord}
			combinedShipmentMap[dest] = combShipment
		}

		if bacthMaker == orderBatchSize {
			for _, comb := range combinedShipmentMap {
				log.Printf("Shipping : %v -> %v", comb.Id, len(comb.OrdersList))
				if err := stream.Send(comb); err != nil {
					return err
				}
			}
			bacthMaker = 0
			combinedShipmentMap = make(map[string]*pb.CombinedShipment)
			continue
		}
		bacthMaker++
	}
}
