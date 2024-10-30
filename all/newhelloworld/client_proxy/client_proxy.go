package clientproxy

import (
	"all/newhelloworld/handle"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

// NewhelloServiceClient 创建一个新的 HelloService 客户端
func NewhelloServiceClient(protocol, address string) *HelloServiceStub {
	conn, err := rpc.Dial(protocol, address)
	if err != nil {
		panic("connect error: " + err.Error())
	}
	return &HelloServiceStub{conn}
}

// hello 调用远程的 Hello 方法
func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handle.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
