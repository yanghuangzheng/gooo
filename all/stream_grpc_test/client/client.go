package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"all/stream_grpc_test/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	//服务端流模式
	res, err := c.GetStream(context.Background(), &proto.StreamReqData{Data: "mooc"})
	if err != nil {
		panic(err)
	}
	for {
		a, err := res.Recv()
		if err != nil {
			break
			//panic(err)
		}
		fmt.Println(a)
	}

	//客户端流模式
	puts, _ := c.PostStream(context.Background())
	i := 0
	for {
		i++
		puts.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("mooc%d", i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	//双端流模式
	allstr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allstr.Recv()
			fmt.Println("收到客户端消息" + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			allstr.Send(&proto.StreamReqData{Data: "我是客户端"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()

}
