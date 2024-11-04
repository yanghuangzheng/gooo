package main

import (
	"context"
	"fmt"
	"time"

	"all/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	//"google.golang.org/grpc/codes"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = c.SayHello(ctx, &proto.HelloRequest{Name: "bobby"})
	if err != nil {
		//panic(err)
		st, ok := status.FromError(err)
		if !ok {
			panic("解析error错误")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}
	//fmt.Println(r.Message)
}
