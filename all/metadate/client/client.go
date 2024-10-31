package main

import (
	"context"
	"fmt"

	//"time"

	"all/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	// md :=metadata.Pairs("timestamp",time.Now().Format(timestampFormat))
	md := metadata.New(map[string]string{
		"name":     "bobby",
		"pasworld": "imooc",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.SayHello( /*context.Background()*/ ctx, &proto.HelloRequest{Name: "bobby"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
