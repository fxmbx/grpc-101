package main

import (
	"context"
	"fmt"
	"log"

	calculator "github.com/fxmbx/grpc-101/calculator/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Calculator clientℹ️")
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("😞 could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	grpcClient := calculator.NewCalculatorServiceClient(conn)
	doUnary(grpcClient)
}

func doUnary(grpcClient calculator.CalculatorServiceClient) {
	fmt.Println("starting uniary rpc")
	request := &calculator.SumRequest{
		FirstNumer:   10,
		SecondNumber: 20,
	}
	response, err := grpcClient.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("😞 error calling SUM RPC: %v", err)
	}

	log.Printf("Response from SUM: %d", response.Result)
}
