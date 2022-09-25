package main

import (
	"fmt"
	"log"
	"net"

	calculator "github.com/fxmbx/grpc-101/calculator/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	calculator.UnimplementedCalculatorServiceServer
}

func main() {
	fmt.Println("calculator 🚀🚀")
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("😞 something went wrong : %v", err)
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(grpcServer, &Server{})

	fmt.Println("➿ grpc CALCULATOR server listening on port 50052")
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("😞 something went wrong getting grpc server to listen: %v", err)
	}
}
