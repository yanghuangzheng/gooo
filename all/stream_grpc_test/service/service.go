package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"all/stream_grpc_test/proto"

	"google.golang.org/grpc"
)

const PORT = "50052"

type server struct {
}

// 服务端流模式没有context
func (s *server) GetStream(in *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		if err := res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}
func (s *server) PostStream(stream proto.Greeter_PostStreamServer) error {
	for {
		if a, err := stream.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}
func (s *server) AllStream(stream proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := stream.Recv()
			fmt.Println("收到客户端消息" + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			stream.Send(&proto.StreamResData{Data: "我是服务器"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}
func main() {
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		panic("ccc")
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic("failed to serve: " + err.Error())
	}
}
