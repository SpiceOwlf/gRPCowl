package main

import (
	"context"
	"log"
	"net"

	pb "github.com/grpcowl/og/helloworld"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (out *pb.HelloReply, err error) {
	log.Printf("Received input: %v %v", in.GetName(), in.GetNickname())
	return &pb.HelloReply{Message: "Heeeeeello " + in.Name + "   " + in.Nickname}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})
	log.Printf("server is up")
	grpcServer.Serve(lis)

}
