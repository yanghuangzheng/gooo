package main

import (
	"context"
	"net"

	//"fmt"

	"all/validate/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	//proto.UnimplementedGreeterServer
	//mustEmbedUnimplementedGreeterServer(),
}

// ctx主要解决协程超时
func (s *Server) SayHello(ctx context.Context, in *proto.Person) (*proto.Person, error) {
	return &proto.Person{
		Id:     32,
		Mobile: "1888888888888888888",
	}, nil
}

type Validator interface {
	Validate() error
}

func main() {
	/*p:=new(proto.Person)
		err :=p.Validate()
	     if err!= nil{
			 panic(err)
		 }*/
	inter := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if r, ok := req.(Validator); ok { //注意上文的接口
			if err := r.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		return handler(ctx, req)
	}
	//生成拦截器
	opt := grpc.UnaryInterceptor(inter)
	g := grpc.NewServer(opt)
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
}
