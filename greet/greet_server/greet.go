package main

import (
	"context"

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
