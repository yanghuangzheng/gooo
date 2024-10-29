package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kirinlabs/HttpRequest" // 请确保这个库存在且正确安装
)

type Response struct {
	Data int `json:"data"`
}

// rpc远程过程调用如何做到像本地调用
func main() {
	println(add(1, 2))
}

func add(a, b int) int {
	req := HttpRequest.NewRequest()
	res, err := req.Get(fmt.Sprintf("http://127.0.0.1:8000/add?a=%d&b=%d", a, b))
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	body, err := res.Body()
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	fmt.Println(string(body)) // 打印响应体以调试

	var rsp Response
	if err := json.Unmarshal(body, &rsp); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	return rsp.Data
}
