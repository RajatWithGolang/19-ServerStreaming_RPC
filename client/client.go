package main

import (
	"context"
	"fmt"
	"io"
	"log"

	greetpb "github.com/Rajat2019/GRPC_IN_ACTION/02-ServerStreaming/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetManyServiceClient(conn)

	doServerStreaming(c)

}

func doServerStreaming(c greetpb.GreetManyServiceClient) {
	fmt.Println("Starting GreetManyTime RPC...")
	req := &greetpb.GreetManyRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rajat",
			LastName:  "Rawat",
		},
	}
	//fmt.Printf("Created Client %f", c)
	res, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling rpc: %v", err)
	}
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}

}
