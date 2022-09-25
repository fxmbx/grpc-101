package main

import (
	"context"
	"log"

	calculator "github.com/fxmbx/grpc-101/calculator/pb"
)

func (*Server) Sum(ctx context.Context, req *calculator.SumRequest) (*calculator.SumResponse, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")

	res := &calculator.SumResponse{
		Result: req.GetFirstNumer() + req.GetSecondNumber(),
	}
	return res, nil
}

func (*Server) PrimeNumberDecomposer(req *calculator.PrimeNumberDecomposerRequest, stream calculator.CalculatorService_PrimeNumberDecomposerServer) error {
	log.Println("Prime number decomposer stream")
	K := req.GetK()
	N := req.GetNumber()
	for N > 1 {
		if N%K == 0 {
			response := &calculator.PrimeNumberDecomposeResponse{
				Result: K,
			}
			stream.Send(response)
			N = N / K
		} else {

			K += 1
		}
	}
	return nil
}
