package main

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
	"gitlab.com/calculator-service/proto"
	"google.golang.org/grpc"
	"io"
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

func (*server) AverageSvc(stream proto.CalculatorService_AverageSvcServer) error {
	sum := 0.0
	count := 0.0
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.Average{
				Num: sum / count,
			})
		}
		if err != nil {
			log.Errorf("error in receiving msg %v", err)
		}
		sum += float64(msg.Num)
		count += 1
	}
	stream.Recv()
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
