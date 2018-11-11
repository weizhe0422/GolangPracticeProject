package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/weizhe0422/GolangPractice/FromUdemy/gRPC_Pracs/gRPC_Greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v", req)
	firstname := req.GetGreeting().GetFirstname()
	result := "Hello," + firstname
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GretManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GretManyTimesServer) error {
	log.Printf("GretManyTimes function was invoked with %v", req)
	firstname := req.GetGreeting().GetFirstname()
	for i := 0; i <= 10; i++ {
		result := "Hello " + firstname + " number " + strconv.Itoa(i)
		res := &greetpb.GreeetManyTimesResoinse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(time.Second)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked with streaming request")
	result := ""
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while receive stream request: %v", err)
		}

		firstName := request.GetGreeting().GetFirstname()
		result += "Hello, " + firstName + " !"

	}
}

func main() {
	fmt.Println("Hello world!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
