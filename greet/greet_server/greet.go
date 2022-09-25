package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	greet "github.com/fxmbx/grpc-101/greet/pb"
)

func (server *Server) Greet(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	result := "Hello " + firstName + "🕺" + lastName
	res := &greet.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*Server) GreetManyTimes(req *greet.GreetManyTimesRequest, stream greet.GreetService_GreetManyTimesServer) error {
	fmt.Println("Greet many stream ➰")
	firstname := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstname + "number " + strconv.Itoa(i)
		res := &greet.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(time.Second)
	}
	return nil
}
