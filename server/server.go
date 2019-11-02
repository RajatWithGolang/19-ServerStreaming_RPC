package main

import (
	"log"
	"net"
	"strconv"
	"time"

	greetpb "github.com/Rajat2019/GRPC_IN_ACTION/02-ServerStreaming/proto"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetManyTimes(req *greetpb.GreetManyRequest, stream greetpb.GreetManyService_GreetManyTimesServer) error {
	first_name := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + first_name + " this is " + strconv.Itoa(i+1) + " response"
		res := &greetpb.GreetManyResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("this is an error %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetManyServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
