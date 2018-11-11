package main

import (
	"context"
	"fmt"
	"github.com/weizhe0422/GolangPractice/FromUdemy/gRPC_Pracs/gRPC_Greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {

	fmt.Println("Hi, I'm client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	//fmt.Printf("Create a client %f", c)

	//doUnary(c)
	//doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			Firstname: "WeiZhe",
			Lastname:  "Chang",
		},
	}
	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("erro while calling Greeting rpc: %v", err)
	}
	log.Printf("Response from Greeting: %v", resp.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			Firstname: "Ray",
			Lastname:  "Chang",
		},
	}
	reqStream, err := c.GretManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error when calling Server streaming: %v", err)
	}

	for {
		msg, err := reqStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error when receive message: %v", err)
		}

		fmt.Printf("Response from GretManyTimes: %v\n", msg.GetResult())
	}

}

func doClientStreaming(c greetpb.GreetServiceClient) {
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error to calling LongGreet: %v", err)
	}

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Ray",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "John",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Kyle",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Scott",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Kobe",
			},
		},
	}

	for _, req := range requests {
		fmt.Printf("Sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Microsecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreeting: %v", err)
	}
	fmt.Printf("Long Greeting response: %v\n", res)
}
