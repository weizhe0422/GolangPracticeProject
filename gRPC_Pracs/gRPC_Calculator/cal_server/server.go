package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"

	"github.com/weizhe0422/GolangPractice/FromUdemy/gRPC_Pracs/gRPC_Calculator/calculatorpb"
)

type server struct{}

func (*server) ADD(ctx context.Context, req *calculatorpb.AddReauest) (*calculatorpb.CalResult, error) {
	fmt.Printf("Add function was invoked with %v", req)
	firstNum := req.GetFirstNum()
	secondNum := req.GetSecondNum()
	res := &calculatorpb.CalResult{
		Calresult: firstNum + secondNum,
	}

	return res, nil
}

func (*server) AVG(stream calculatorpb.AddService_AVGServer) error {
	temp := int32(0)
	inputCnt := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result := float64(temp) / float64(inputCnt)
			return stream.SendAndClose(&calculatorpb.AvgResult{
				CalResule: result,
			})
		}
		if err != nil {
			log.Fatalf("error while reveive request: %v", err)
		}
		temp = temp + req.GetNumber()
		inputCnt++
	}
	return nil
}

func main() {
	fmt.Println("This is from server side")

	lis, err := net.Listen("tcp", "0.0.0.0:50021")
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterAddServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Printf("failed when grpc serve: %v", err)
	}

}
