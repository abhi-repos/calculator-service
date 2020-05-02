package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gitlab.com/calculator-service/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("Error in connection %v", err)
	}
	defer conn.Close()
	cc := proto.NewCalculatorServiceClient(conn)
	sum, err := GetSum(10, 3, cc)
	if err != nil {
		log.Printf("Error in getting response", err)
	}
	fmt.Println("SUM OF 10 and 3 is", sum)
	mult, err := GetProduct(10, 3, cc)
	if err != nil {
		log.Printf("Error in getting response", err)
	}
	fmt.Println("Multiplication OF 10 and 3 is", mult)
}

func GetSum(first, second int64, client proto.CalculatorServiceClient) (int64, error) {
	req := &proto.Request{
		Input: &proto.Input{
			FirstNumber:  first,
			SecondNumber: second,
		},
	}
	res, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error in getting response", err)
	}
	return res.Result, nil
}

func GetProduct(first, second int64, client proto.CalculatorServiceClient) (int64, error) {
	req := &proto.Request{
		Input: &proto.Input{
			FirstNumber:  first,
			SecondNumber: second,
		},
	}
	res, err := client.Multiply(context.Background(), req)
	if err != nil {
		log.Fatalf("Error in getting response", err)
	}
	return res.Result, nil
}
