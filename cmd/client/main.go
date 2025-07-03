package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/gabrielfeb/list-orders-challenge-go/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	// Test gRPC call
	res, err := client.ListOrders(context.Background(), &pb.Blank{})
	if err != nil {
		log.Fatalf("Error calling ListOrders: %v", err)
	}

	log.Printf("Orders: %v", res.GetOrders())
}
