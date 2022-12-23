package main

import (
	"context"
	"fmt"

	pb "github.com/udholdenhed/grpc-study/greet/proto"
)

type Server struct {
	pb.GreetServiceServer
}

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	return &pb.GreetResponse{
		Result: fmt.Sprintf("Hello %s!", in.FirstName),
	}, nil
}
