package main

import (
	"context"
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
	doUnary(grpcClient)
}

func doUnary(grpcClient greet.GreetServiceClient) {
	fmt.Println("starting uniary rpc")
	request := &greet.GreetRequest{
		Greeting: &greet.Greeting{
			FirstName: "funbi",
			LastName:  "olaore",
		},
	}
	res, err := grpcClient.Greet(context.Background(), request)
	if err != nil {
		log.Fatalf("😞 error calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}
