package handle

// 为了解决名称冲突
const HelloServiceName = "handle/HelloService"

//我们关心的是HelloServer还是这个结构体的方法
type HelloServer struct{}

func (s *HelloServer) Hello(req string, reply *string) error {
	*reply = "hello, " + req
	return nil
}

//这些概念在grpc都有对应
//server——proxy与client——proxy能否自动生成 为多种语言生成
//3.都能满足 这个就是protobuf+grpc
