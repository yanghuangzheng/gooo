package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloServer struct{}

func (s *HelloServer) Hello(req string, reply *string) error {
	*reply = "hello, " + req
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	err = rpc.RegisterName("HelloService", new(HelloServer))
	if err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go rpc.ServeConn(conn)
		//一连串的代码大部分好像都是net的包好像和rpc没有用关系
		//不行 rpc调用中有几个问题需要解决1.call id2 序列化与反序列化
		//可以跨语言调用呢 1.go语言的序列化与反序列化的协议是什么（gob）2.能否能替换成常见的序列化
	}
}
