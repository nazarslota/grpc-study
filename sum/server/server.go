package main

import (
	"context"

	pb "github.com/nazarslota/grpc-study/sum/proto"
)

type Server struct {
	pb.SumServiceServer
}

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	return &pb.SumResponse{Result: in.A + in.B}, nil
}
