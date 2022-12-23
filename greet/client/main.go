package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "5051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) { _ = conn.Close() }(conn)

}
