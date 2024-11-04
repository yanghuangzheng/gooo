package main

import (
	"context"
	"net"
	"time"

	//"time"

	"all/proto"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/status"
	//"google.golang.org/grpc/codes"
)

type Server struct {
	//proto.UnimplementedGreeterServer
}

// ctx主要解决协程超时
func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	//返回错误代码
	//return nil,status.Errorf(codes.NotFound,"记录未找到：%s",in.Name)
	time.Sleep(5 * time.Second)
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
	g.Serve(lis)

}
