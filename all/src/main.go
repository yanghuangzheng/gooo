package main

import (
	"fmt"
)

func add(a, b int) int {
	total := a + b
	return total
}

type Company struct {
	Name    string
	Address string
}
type Employee struct {
	Name    string
	company Company
}
type Pringresult struct {
	In  string
	Err error
}

func rpc(employee Employee) {
	/*支持远程调用的函数
		1.建立连接 tcp/http
		2.将employee对象序列化成json字符串-序列化
		3.发送json字符串 - 调用成功好实际接收到的是一个二进制的数据
		4.等待服务器发送结果
		5.将服务器返回的数据解析成Pritresult
		1.监听网络短裤
		2.读取数据-对数据进行反序列化employee对象
		3.开始处理业务逻辑
		4.将出来的结果printresult序列话json二进制数据
		5.返回数据
		序列化和反序列化是可以选择的，不一定json
	     rpc第二点
		  http协议有一个问题:一次性 一但对方返回结果 连接断开 http2.0 长连接grpc
	*/

}

func main() {
	//想把add变成一个远程函数调用，吧add放在远程服务器上
	/*
		例如扣减库存 将库存服务放在远程服务器上调用 那如何调用
		一定牵扯网络，做成一个web服务（gin，beego，net、httpserver）
		1.函数的调用参数如何传递-json（json是一种数据格式的协议） 、xml、protobuf,mspage-编码协议 json并不是一个高性能的编码协议
		网络调用的两个端 服务端-》gin负责解析数据  客户端-》将数据传输到gin
	*/
	//将打印的工资放在另一台服务器上，我需要将本地的内存对象//可行方式就是讲struct对象序列为json对象 同时远程服务器反解对象
	fmt.Println(add(1, 2))
	fmt.Println(Employee{
		Name: "bobby",
		company: Company{
			Name:    "mooc",
			Address: "北京",
		},
	})
}
