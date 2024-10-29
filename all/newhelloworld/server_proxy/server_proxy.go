package serverproxy

import (
	"good/newhelloworld/handle"
	"net/rpc"
)

type HelloServer interface {
	Hello(req string, reply *string) error
}

// 如何做到解耦-我们关心的是函数 鸭子类型
func RegisterHelloServer(srv HelloServer) error {
	return rpc.RegisterName(handle.HelloServiceName, srv)
}
