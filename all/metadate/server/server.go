package main

import (
	"context"
	"fmt"
	"net"

	"all/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	//proto.UnimplementedGreeterServer
}

// ctx主要解决协程超时
func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("get data err")
	}
	if name, ok := md["name"]; ok {
		fmt.Println(name)
		for i, e := range name {
			fmt.Println(i, e)
		}
	}
	for key, val := range md {
		fmt.Println(key, val)
	}
	return &proto.HelloReply{
		Message: "hello hhh" + in.Name,
	}, nil
}
func main() {
	g := grpc.NewServer()
	//var serverGreeter = &Server{}
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
}
