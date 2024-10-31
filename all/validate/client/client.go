package main

import (
	"context"
	"fmt"

	//"time"
	"all/validate/proto"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/metadata"
)

type custom struct {
}

func (c custom) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "100",
		"appkey": "keey",
	}, nil
}
func (c custom) RequireTransportSecurity() bool {
	return false
}
func main() {
	/*inter := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	                   start := time.Now()
					   //逻辑
					  md :=metadata.New(map[string]string{
						"appid":"100",
						"appkey":"keey",
					})
					//ctxx:=metadata.NewOutgoingContext(context.Background(),md)
					   err:=invoker(ctxx,method,req,reply,cc,opts...)
					   fmt.Println("耗时：%s\n",time.Since(start))
					   return err
	   }*/
	opt := grpc.WithPerRPCCredentials(custom{})
	//opt:=grpc.WithUnaryInterceptor(inter)
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
	r, err := c.SayHello(context.Background(), &proto.Person{
		Id:     1000,
		Email:  "bobby@imocc.com",
		Mobile: "18888888888",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Mobile)
}
