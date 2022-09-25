package main

import (
	"context"
	"fmt"
	"io"
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
	doServerStreaming(grpcClient)
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

func doServerStreaming(grpcClient greet.GreetServiceClient) {
	fmt.Println("starting server streaming  rpc")

	request := &greet.GreetManyTimesRequest{
		Greeting: &greet.Greeting{
			FirstName: "John",
			LastName:  "Snow",
		},
	}
	responseStream, err := grpcClient.GreetManyTimes(context.Background(), request)
	if err != nil {
		log.Fatalf("😞 error calling GreetMany Times RPC: %v", err)
	}
	for {
		msg, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("😞 error whilig reading stream: %v", err)

		}
		log.Printf("Response from GreetManyTime: %v", msg.GetResult())
	}

}
