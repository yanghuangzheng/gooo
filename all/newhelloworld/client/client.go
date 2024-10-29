package main

import (
	"fmt"
	"log"

	clientproxy "good/newhelloworld/client_proxy"
)

func main() {
	// 建立连接
	client := clientproxy.NewhelloServiceClient("tcp", "localhost:1234")
	defer client.Close()
	//只写业务层不管
	//客户端部分
	var rp string
	err := client.Hello("boby", &rp)
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}
	fmt.Println(rp)
}
