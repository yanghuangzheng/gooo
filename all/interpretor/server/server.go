package main

import (
	"context"
	"fmt"
	"net"

	"all/proto"

	"google.golang.org/grpc"
)

type Server struct {
	//proto.UnimplementedGreeterServer
}

// ctx主要解决协程超时
func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello hhh" + in.Name,
	}, nil
}
func main() {
	inter := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到了一个新请求")
		res, err := handler(ctx, req)
		fmt.Println("请求已经完成")
		return res, err
		// return handler(ctx,req)
	}
	//生成拦截器
	opt := grpc.UnaryInterceptor(inter) //grpc生成
	g := grpc.NewServer(opt)            //配置拦截器可以传入多个opt
	//var serverGreeter = &Server{}
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
}
