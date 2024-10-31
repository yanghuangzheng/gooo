package all
//protobuf的底层实现
//https://github.com/protocolbuffers/protobuf/releases
protoc -I . goods.proto --go_out=plugins=grpc:.
protoc -I . helloworld.proto --go_out=plugins=grpc:.
protoc -I . helloworld.proto --go_out=. --go-grpc_out=.
protoc -I . --go_out=. --go-grpc_out=. --validate_out="lang=go:." form.proto
"google.golang.org/grpc/metadata"
go-grpc-middleware
"google.golang.org/grpc/status"
google.golang.org/grpc/codes
/////////////////////////////////////////////
Validate
/////////////////////////////////////////////
form.proto
syntax="proto3";

import "validate.proto";
option go_package = ".;proto";

service Greeter{
    rpc SayHello(Person) returns(Person);
}
//好比文档 表单验证 
//验证器
message Person{
    uint64 id=1 [(validate.rules).uint64.gt   =999];
    string email =2 [(validate.rules).string.email   =true];
    string mobile = 3 [(validate.rules).string = {
        pattern:   "^1[3456789]\\d{9}$",
        max_bytes: 256,
      }];
}
////////////////////////////////////////////////
service
///////////////////////////////////////////////
package main
import(
	"context"
	"net"
	//"fmt"

	"all/Validator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
type Server struct{
	//proto.UnimplementedGreeterServer
	//mustEmbedUnimplementedGreeterServer(),
}
//ctx主要解决协程超时
func (s *Server) SayHello(ctx context.Context,in*proto.Person) (*proto.Person, error) {
    return &proto.Person{
       Id:32,
	   Mobile:"1888888888888888888",
    }, nil
}
type Validator interface{
	Validate() error
}
func main(){
	/*p:=new(proto.Person)
	err :=p.Validate()
     if err!= nil{
		 panic(err)
	 }*/
	inter := func(ctx context.Context,req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler)(resp interface{},err error){
		if r,ok:=req.(Validator);ok{//注意上文的接口
			if err:=r.Validate();err!=nil{
				return nil,status.Error(codes.InvalidArgument,err.Error())
			}
		}
          return handler(ctx,req)
	}
	//生成拦截器
	opt:=grpc.UnaryInterceptor(inter)
	g:=grpc.NewServer(opt)
	proto.RegisterGreeterServer(g,&Server{})
	lis,err:=net.Listen("tcp","0.0.0.0:8080")
	if err!=nil {
		panic("failed to listen:"+err.Error())
	}
	err=g.Serve(lis)
}
/////////////////////////////////////////////////////////////
client
////////////////////////////////////////////////////////////////
package main
import(
	"context"
	"fmt"
	//"time"

	"all/Validator/proto"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/metadata"
)
type custom struct{
}
func(c custom)GetRequestMetadata(ctx context.Context,uri ...string)(map[string]string,error) {
	return map[string]string{
		"appid":"100",
		"appkey":"keey",
	},nil
}
func (c custom)RequireTransportSecurity()bool{
	return false
}
func main(){
	/*inter := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
                   start := time.Now()
				   //逻辑
				  md :=metadata.New(map[string]string{
					"appid":"100",
					"appkey":"keey",
				})
				//ctxx:=metadata.NewOutgoingContext(context.Background(),md)
				   err:=invoker(ctxx,method,req,reply,cc,opts...)
				   fmt.Println("耗时：%s\n",time.Since(start))
				   return err
   }*/
    opt:=grpc.WithPerRPCCredentials(custom{})
	//opt:=grpc.WithUnaryInterceptor(inter)
	//var opts[]grpc.DialOption
	//opts=append(opts,grpc.WithInsecure())
	//opts=append(grpc.WithUnaryInterceptor(inter)
	//conn,err:=grpc.Dial("127.0.0.1:8080",opts...)
	 conn,err:=grpc.Dial("127.0.0.1:8080",grpc.WithInsecure(),opt)
	 if err!=nil{
		 panic(err)
	 }
	 defer conn.Close()
	 
	 c:=proto.NewGreeterClient(conn)
	 r,err:=c.SayHello(context.Background(),&proto.Person{
		Id:1000,
		Email:"bobby@imocc.com",
	    Mobile:"18888888888",
	})
	 if err!=nil{
		panic(err)
	}
	fmt.Println(r.Mobile)
}



////////////////////////////////////////////////////////////////
interpretor
///////////////////////////////////////////////////////////////
service
//////////////////////////////////////////////////////////////////
package main
import(
	"context"
	"net"
	"fmt"

	"all/grpc_test/proto"

	"google.golang.org/grpc"
)
type Server struct{
	//proto.UnimplementedGreeterServer
}
//ctx主要解决协程超时
func (s *Server) SayHello(ctx context.Context,in*proto.HelloRequest) (*proto.HelloReply, error) {
    return &proto.HelloReply{
        Message :"hello hhh" + in.Name,
    }, nil
}
func main(){
	inter := func(ctx context.Context,req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler)(resp interface{},err error){
	     fmt.Println("接收到了一个新请求")
		 res,err:=handler(ctx,req)
		 fmt.Println("请求已经完成")
		 return res,err
         // return handler(ctx,req)
	}
	//生成拦截器
	opt:=grpc.UnaryInterceptor(inter)//grpc生成
	g:=grpc.NewServer(opt)//配置拦截器可以传入多个opt
	//var serverGreeter = &Server{}
	proto.RegisterGreeterServer(g,&Server{})
	lis,err:=net.Listen("tcp","0.0.0.0:8080")
	if err!=nil {
		panic("failed to listen:"+err.Error())
	}
	err=g.Serve(lis)
}
/////////////////////////////////////////////////////////////////////////////////
client
/////////////////////////////////////////////////////////////////////////////
package main
import(
	"context"
	"fmt"
	"time"

	"all/grpc_test/proto"

	"google.golang.org/grpc"
)
func main(){
	inter := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
                   start := time.Now()
				   //逻辑
				   err:=invoker(ctx,method,req,reply,cc,opts...)
				   fmt.Println("耗时：%s\n",time.Since(start))
				   return err
   }
	opt:=grpc.WithUnaryInterceptor(inter)
	//var opts[]grpc.DialOption
	//opts=append(opts,grpc.WithInsecure())
	//opts=append(grpc.WithUnaryInterceptor(inter)
	//conn,err:=grpc.Dial("127.0.0.1:8080",opts...)
	 conn,err:=grpc.Dial("127.0.0.1:8080",grpc.WithInsecure(),opt)
	 if err!=nil{
		 panic(err)
	 }
	 defer conn.Close()
	 
	 c:=proto.NewGreeterClient(conn)
	 r,err:=c.SayHello(context.Background(),&proto.HelloRequest{Name:"bobby"})
	 if err!=nil{
		panic(err)
	}
	fmt.Println(r.Message)
}
////////////////////////////////////////////////////////////////////////////////////////////



metadate
/////////////////////////
service
///////////////////////////////
package main
import(
	"context"
	"net"
	"fmt"

	"all/grpc_test/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)
type Server struct{
	//proto.UnimplementedGreeterServer
}
//ctx主要解决协程超时
func (s *Server) SayHello(ctx context.Context,in*proto.HelloRequest) (*proto.HelloReply, error) {
	md,ok:=metadata.FromIncomingContext(ctx)
	if ok{
		fmt.Println("get data err")
	}
	if name,ok:=md["name"];ok{
		fmt.Println(name)
		for i,e:=range name{
			fmt.Println(i,e)
		}
	}
        for key,val:=range md{
			fmt.Println(key,val)}
    return &proto.HelloReply{
        Message :"hello hhh" + in.Name,
    }, nil
}
func main(){
	g:=grpc.NewServer()
	//var serverGreeter = &Server{}
	proto.RegisterGreeterServer(g,&Server{})
	lis,err:=net.Listen("tcp","0.0.0.0:8080")
	if err!=nil {
		panic("failed to listen:"+err.Error())
	}
	err=g.Serve(lis)
}
//////////////////////////////////////////////////////////////////////////////////
client
/////////////////////////////////////////////////////////////////
package main
import(
	"context"
	"fmt"
	//"time"

	"all/grpc_test/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)
func main(){
	 conn,err:=grpc.Dial("127.0.0.1:8080",grpc.WithInsecure())
	 if err!=nil{
		 panic(err)
	 }
	 defer conn.Close()
	 
	 c:=proto.NewGreeterClient(conn)
	// md :=metadata.Pairs("timestamp",time.Now().Format(timestampFormat))
	 md :=metadata.New(map[string]string{
		 "name":"bobby",
		 "pasworld":"imooc",
	 })
	 ctx:=metadata.NewOutgoingContext(context.Background(),md)
	 r,err:=c.SayHello(/*context.Background()*/ctx,&proto.HelloRequest{Name:"bobby"})
	 if err!=nil{
		panic(err)
	}
	fmt.Println(r.Message)
}
、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、
syntax="proto3";
option go_package = ".;proto";
service Greeter{
    rpc SayHello(HelloRequest) returns(HelloReply);
}
message HelloRequest{
    string Name =1;
}
message HelloReply{
    string message=1;
}

