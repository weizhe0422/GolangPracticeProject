package main

import (
	"context"
	"fmt"
	"github.com/weizhe0422/GolangPractice/FromUdemy/gRPC_Pracs/gRPC_Calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

func main() {
	fmt.Println("This is from client side")

	conn, err := grpc.Dial("0.0.0.0:50021", grpc.WithInsecure())
	if err != nil {
		log.Printf("fail to connect: %v", err)
	}

	defer conn.Close()

	c := calculatorpb.NewAddServiceClient(conn)

	//doUnary(c)
	doClientStreaming(c)
}

func doUnary(c calculatorpb.AddServiceClient) {
	req := &calculatorpb.AddReauest{
		FirstNum:  12345,
		SecondNum: 67890,
	}
	result, err := c.ADD(context.Background(), req)
	if err != nil {
		log.Printf("error while add two numbers: %v\n", err)
	}

	log.Printf("Respons from ADD: %v", result.Calresult)
}

func doClientStreaming(c calculatorpb.AddServiceClient) {

	stream, err := c.AVG(context.Background())
	requests := []*calculatorpb.AvgRequest{
		&calculatorpb.AvgRequest{
			Number: 1,
		},
		&calculatorpb.AvgRequest{
			Number: 2,
		},
		&calculatorpb.AvgRequest{
			Number: 3,
		},
		&calculatorpb.AvgRequest{
			Number: 4,
		},
	}

	for idx, req := range requests {
		fmt.Printf("The %v number is %v\n", strconv.Itoa(idx), req)
		stream.Send(req)
		time.Sleep(1000000 * time.Microsecond)
	}
	time.Sleep(1000000 * time.Microsecond)
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receive Average response: %v", err)
	}
	fmt.Printf("The result of AVG response: %v", response.CalResule)
}
