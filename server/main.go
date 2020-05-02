package main

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
	"gitlab.com/calculator-service/proto"
	"google.golang.org/grpc"
	"net"
)

const address = "0.0.0.0:50051"

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

func (*server) PrimeDecomposition(req *proto.RequestDecomposition, stream proto.CalculatorService_PrimeDecompositionServer) error {
	num := int(req.Input.Number)
	k := 2
	for {
		if num <= 1 {
			break
		}
		if num%k == 0 {
			err := stream.Send(&proto.ResponseDecomposition{
				PrimeFactor: int64(k),
			})
			if err != nil {
				log.Fatalf("Error while sending %v", err)
			}
			num /= k
		} else {
			k = k + 1
		}

	}
	return nil
}

func main() {
	fmt.Println("server initialization")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("error in listenting %v", err)
	}
	var svc server
	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &svc)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("error in listening %v", err)
	}
}
