package main

import (
	"log"
	"net"
	"net/rpc"

	"good/newhelloworld/handle"
	serverproxy "good/newhelloworld/server_proxy"
)

type HelloServer struct{}

func (s *HelloServer) Hello(req string, reply *string) error {
	*reply = "hello, " + req
	return nil
}

func main() {
	//实例化一个server
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()
	// 注册处理逻辑handler
	err = serverproxy.RegisterHelloServer(&handle.HelloServer{})
	//err = rpc.RegisterName(handle.HelloServiceName, &handle.HelloServer{})
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
	}
}
