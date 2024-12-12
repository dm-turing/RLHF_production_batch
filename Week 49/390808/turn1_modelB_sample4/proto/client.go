package main

import (
	"context"
	"log"

	pb "github.com/example/proto"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer cc.Close()

	c := pb.NewHelloClient(cc)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "you"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r.Message)
}
