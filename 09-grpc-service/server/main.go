package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-demo/proto" 

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCurrencyConverterServer
}

func (s *server) GetExchangeRate(ctx context.Context, req *pb.RateRequest) (*pb.RateResponse, error) {
	fmt.Printf("Received request: %v -> %v\n", req.FromCurrency, req.ToCurrency)

	rates := map[string]float32{
		"USD_EUR": 0.92,
		"USD_IRR": 600000.0,
		"EUR_USD": 1.09,
	}

	key := fmt.Sprintf("%s_%s", req.FromCurrency, req.ToCurrency)
	val, exists := rates[key]

	if !exists {
		return nil, fmt.Errorf("currency pair not found")
	}

	return &pb.RateResponse{Rate: val}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	
	pb.RegisterCurrencyConverterServer(s, &server{})

	fmt.Println("gRPC Server started at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}