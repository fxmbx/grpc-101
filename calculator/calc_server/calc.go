package main

import (
	"context"

	calculator "github.com/fxmbx/grpc-101/calculator/pb"
)

func (*Server) Sum(ctx context.Context, req *calculator.SumRequest) (*calculator.SumResponse, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")

	res := &calculator.SumResponse{
		Result: req.GetFirstNumer() + req.GetSecondNumber(),
	}
	return res, nil
}
