package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 建立连接
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("连接失败: " + err.Error())
	}
	defer client.Close()

	var rp string                                        // 使用字符串类型而不是指针
	err = client.Call("HelloService.Hello", "boby", &rp) // 注意这里使用&rp
	//确定我们必须知道name是什么
	// 比如我们想调用的是client。hello（“bobby”，&reply）
	if err != nil {
		panic("调用失败: " + err.Error())
	}
	fmt.Println(rp)
}
