package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCurrencyConverterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.GetExchangeRate(ctx, &pb.RateRequest{
		FromCurrency: "USD",
		ToCurrency:   "IRR",
	})

	if err != nil {
		log.Fatalf("Error calling server: %v", err)
	}

	fmt.Printf("Rate USD to IRR: %.0f\n", res.Rate)
}