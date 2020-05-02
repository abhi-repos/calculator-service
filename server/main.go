package main

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
	"gitlab.com/calculator-service/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	sum := req.Input.FirstNumber + req.Input.SecondNumber
	res := &proto.Response{
		Result: sum,
	}
	return res, nil
}
func (*server) Multiply(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	mult := req.Input.FirstNumber * req.Input.SecondNumber
	res := &proto.Response{
		Result: mult,
	}
	return res, nil
}

func main() {
	fmt.Println("SERVER")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error in listenting %v", err)
	}
	var svc server
	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &svc)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("error in listening %v", err)
	}
}
