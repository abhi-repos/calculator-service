package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gitlab.com/calculator-service/proto"
	"google.golang.org/grpc"
	"io"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Errorf("Error in connection %v", err)
	}
	defer conn.Close()
	cc := proto.NewCalculatorServiceClient(conn)

	sum, err := GetSum(10, 3, cc)
	if err != nil {
		log.Errorf("Error in getting response", err)
	}
	log.Println("SUM OF 10 and 3 is", sum)

	mult, err := GetProduct(10, 3, cc)
	if err != nil {
		log.Errorf("Error in getting response", err)
	}
	log.Println("Multiplication OF 10 and 3 is", mult)

	slc, err := GetPrimeStream(120, cc)
	if err != nil {
		log.Errorf("Error in getting response", err)
	}
	log.Printf("Prime Factor of %v - %v", 120, slc)

	avg, err := GetAverage([]int64{1, 2, 3, 4}, cc)
	if err != nil {
		log.Errorf("Error in getting response", err)
	}
	log.Printf("Average is - %v", avg)
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

func GetPrimeStream(num int64, client proto.CalculatorServiceClient) ([]int64, error) {
	req := &proto.RequestDecomposition{
		Input: &proto.InputDecomposition{
			Number: num,
		},
	}
	var slc []int64
	stream, err := client.PrimeDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error in getting stream %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error in getting msg %v", err)
		}
		prime := msg.PrimeFactor
		log.Println(prime)
		slc = append(slc, prime)
	}
	return slc, nil
}

func GetAverage(nums []int64, client proto.CalculatorServiceClient) (float64, error) {
	stream, err := client.AverageSvc(context.Background())
	if err != nil {
		log.Errorf("error in getting stream %v", err)
	}
	for _, num := range nums {
		stream.Send(&proto.Number{
			Num: num,
		})
	}
	rst, err := stream.CloseAndRecv()
	if err != nil {
		log.Errorf("error in getting response %v", err)
	}
	return rst.Num, nil
}
