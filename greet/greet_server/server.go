package main

import (
	"fmt"
	"log"
	"net"

	greet "github.com/fxmbx/grpc-101/greet/pb"
	"google.golang.org/grpc"
)

type Server struct {
	greet.UnimplementedGreetServiceServer
}

func main() {
	fmt.Println("🚀🚀")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("😞 something went wrong : %v", err)
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()

	greet.RegisterGreetServiceServer(grpcServer, &Server{})
	fmt.Println("➿ grpc server listening on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("😞 something went wrong getting grpc server to listen: %v", err)
	}
}
