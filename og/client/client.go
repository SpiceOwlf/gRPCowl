package main

import (
	"context"
	"log"
	"time"

	pb "github.com/grpcowl/og/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error is %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "michael", Nickname: "owl"})
	if err != nil {
		log.Fatalf("error is %v", err)
	}
	log.Printf("server replied %v", resp.GetMessage())
}
