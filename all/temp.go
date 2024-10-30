package all
protobuf的底层实现
https://github.com/protocolbuffers/protobuf/releases
protoc -I . goods.proto --go_out=plugins=grpc:.
protoc -I . helloworld.proto --go_out=plugins=grpc:.
protoc -I . helloworld.proto --go_out=. --go-grpc_out=.
protobuf的原理

//stream_grpc_test
////////////////////////////////////////////////////
client
/////////////////////////////////////////////////////
package main
import(
	"context"
	"fmt"
	"sync"
	"time"

	"all/stream_grpc_test/proto"

	"google.golang.org/grpc"
)
func main(){
	 conn,err:=grpc.Dial("127.0.0.1:50052",grpc.WithInsecure())
	 if err!=nil{
		 panic(err)
	 }
	 defer conn.Close()
	 c:=proto.NewGreeterClient(conn)

	 //服务端流模式
	 res,err:=c.GetStream(context.Background(),&proto.StreamReqData{Data:"mooc"})
	 if err!=nil{
		panic(err)
	}
	for{
		a,err:=res.Recv()
		if err!=nil{
			break
			//panic(err)
		}
		fmt.Println(a)
	}

	//客户端流模式
	puts,_:=c.PostStream(context.Background())
	i :=0
	for {
		i++
		puts.Send(&proto.StreamReqData{
			Data:fmt.Sprintf("mooc%d",i),
		})
		time.Sleep(time.Second)
		if i>10{
			break
		}
	}

	//双端流模式
	allstr,_:=c.AllStream(context.Background())
		wg:=sync.WaitGroup{}
		wg.Add(2)
		go func(){
			defer wg.Done()
			for{
				data,_:=allstr.Recv()
				fmt.Println("收到客户端消息"+data.Data)
			}
		}()
		go func(){
			defer wg.Done()
			for{
				allstr.Send(&proto.StreamReqData{Data:"我是客户端"})
				time.Sleep(time.Second)
			}
		}()
		wg.Wait()

}
/////////////////////////////////////////////////////////
stream.proto
/////////////////////////////////////////////////////////
syntax="proto3";
option go_package = ".;proto";
service Greeter{
    rpc GetStream(StreamReqData) returns(stream StreamResData);//服务端流模式
    rpc PostStream(stream StreamReqData) returns(StreamResData);//客户端流模式
    rpc AllStream(stream StreamReqData) returns(stream StreamResData);//双向模式
}
message StreamReqData{
    string data =1;
}
message StreamResData{
    string data =1;
}
///////////////////////////////////////////////////////
server
///////////////////////////////////////////////////////
package main
	import(
		"fmt"
		"net"
		"time"
		"sync"
	
		"all/stream_grpc_test/proto"
	
		"google.golang.org/grpc"
	)
const PORT="50052"
type server struct{

}
//服务端流模式没有context
func(s *server)GetStream(in *proto.StreamReqData,res proto.Greeter_GetStreamServer)error{
	i := 0
	for {
		i++
		if err := res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
  }
func(s *server)PostStream(stream proto.Greeter_PostStreamServer) error{
	for{
	  if a,err:=stream.Recv();err!=nil{
		  fmt.Println(err)
		  break
	  }else{
		  fmt.Println(a.Data)
	  }
	}
return nil
}
func(s *server) AllStream(stream proto.Greeter_AllStreamServer) error{
	   wg:=sync.WaitGroup{}
	   wg.Add(2)
	   go func(){
		defer wg.Done()
		   for{
			   data,_:=stream.Recv()
			   fmt.Println("收到客户端消息"+data.Data)
		   }
	   }()
	   go func(){
		   defer wg.Done()
		   for{
			   stream.Send(&proto.StreamResData{Data:"我是服务器"})
			   time.Sleep(time.Second)
		   }
	   }()
	   wg.Wait()
	return nil
}
func main(){
    lis, err := net.Listen("tcp", ":"+PORT)
	if err!=nil{
		panic("ccc")
	}
	s:=grpc.NewServer()
	proto.RegisterGreeterServer(s,&server{})
	if err := s.Serve(lis); err != nil {
		panic("failed to serve: " + err.Error())
	}
}
	///////////////////////////////////////////////////////////////////////
