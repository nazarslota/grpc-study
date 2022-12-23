package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/udholdenhed/grpc-study/greet/proto"
)

const addr = "0.0.0.0:5051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) { _ = conn.Close() }(conn)

	client := pb.NewGreetServiceClient(conn)
	DoGreat(client)
}

func DoGreat(client pb.GreetServiceClient) {
	response, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Nazar",
	})
	if err != nil {
		panic(err)
	}
	log.Println(response.Result)
}
