package main

import (
	"fmt"
	"log"

	greet "github.com/fxmbx/grpc-101/greet/pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("🚴🚴")
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("😞 could not connect to grpc server: %v", err)
	}
	defer conn.Close()
	grpcClient := greet.NewGreetServiceClient(conn)
	fmt.Printf("client %f", grpcClient)
}
