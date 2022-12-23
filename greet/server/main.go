package main

import (
	"net"

	pb "github.com/udholdenhed/grpc-study/greet/proto"
	"google.golang.org/grpc"
)

const addr = "0.0.0.0:5051"

type server struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
