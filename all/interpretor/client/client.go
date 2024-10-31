package main

import (
	"context"
	"fmt"
	"time"

	"all/proto"

	"google.golang.org/grpc"
)

func main() {
	inter := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		//逻辑
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println("耗时：%s\n", time.Since(start))
		return err
	}
	opt := grpc.WithUnaryInterceptor(inter)
	//var opts[]grpc.DialOption
	//opts=append(opts,grpc.WithInsecure())
	//opts=append(grpc.WithUnaryInterceptor(inter)
	//conn,err:=grpc.Dial("127.0.0.1:8080",opts...)
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), opt)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "bobby"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
