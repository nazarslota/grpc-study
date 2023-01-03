package main

import (
	"context"
	"fmt"

	pb "github.com/nazarslota/grpc-study/sum/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "0.0.0.0:5051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) { _ = conn.Close() }(conn)

	client := pb.NewSumServiceClient(conn)
	DoSum(client)
}

func DoSum(client pb.SumServiceClient) {
	response, err := client.Sum(context.Background(), &pb.SumRequest{
		A: 10,
		B: 3,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Result)
}
