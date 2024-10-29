package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 建立连接
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("连接失败: " + err.Error())
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn)) //创建了一个使用 JSON 格式进行编解码的 ClientCodec 实例。这个实例会将 RPC 调用的数据序列化为 JSON 格式并通过网络发送，同时也会从网络接收 JSON 格式的数据并反序列化为 Go 语言的数据结构
	var rp string
	err = client.Call("HelloService.Hello", "boby", &rp)
	if err != nil {
		panic("调用失败: " + err.Error())
	}
	fmt.Println(rp)
}
