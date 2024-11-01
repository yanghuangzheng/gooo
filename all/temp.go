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

OK (0)
成功完成。
CANCELLED (1)
操作被取消（通常是客户端取消请求）。
UNKNOWN (2)
未知错误。通常表示发生了意外错误。
INVALID_ARGUMENT (3)
客户端提供了无效的参数。
DEADLINE_EXCEEDED (4)
请求超时。
NOT_FOUND (5)
请求的资源未找到。
ALREADY_EXISTS (6)
请求的资源已存在。
PERMISSION_DENIED (7)
客户端没有足够的权限执行请求的操作。
RESOURCE_EXHAUSTED (8)
资源耗尽（例如，配额或限制）。
FAILED_PRECONDITION (9)
操作失败，因为前提条件不满足。
ABORTED (10)
操作被中止，通常是因为并发操作导致的冲突。
OUT_OF_RANGE (11)
操作失败，因为索引或键超出范围。
UNIMPLEMENTED (12)
服务未实现请求的方法。
INTERNAL (13)
内部错误。通常表示服务器内部发生了错误。
UNAVAILABLE (14)
服务不可用，通常是由于维护或过载。
DATA_LOSS (15)
发生数据丢失或损坏。
UNAUTHENTICATED (16)
请求未通过身份验证。
"context" 的理解与使用
/////////////////////////////////////////////////////////////
orm //屏蔽底层sql语句 将一张表映射成一个类 表中的列映射成类中的一个类对于go中 列可以映射成struct中的类型 但是数据库中的列要具备很好的描述性，但是struct有tag。
对于个人而言不应该去纠结一个选择哪一个orm框架因为orm迁移成本低
sql语言远比orm重要 一定要注意sql 虽然屏蔽但是很重要
github.com/go-gorm/gorm
github.com/facebook/ent
github.com/facebook/sqlx

go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

https://gorm.io/zh_CN/docs/logger.html
/////////////////////////////////////////////////////////
err_resolve
////////////////////////////////////////////////////////
client
///////////////////////////////////////////////////////////
package main
import(
	"context"
	"fmt"

	"all/grpc_test/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	//"google.golang.org/grpc/codes"
)
func main(){
	 conn,err:=grpc.Dial("127.0.0.1:8080",grpc.WithInsecure())
	 if err!=nil{
		 panic(err)
	 }
	 defer conn.Close()
	 
	 c:=proto.NewGreeterClient(conn)
	 _,err=c.SayHello(context.Background(),&proto.HelloRequest{Name:"bobby"})
	 if err!=nil{
		//panic(err)
		st,ok := status.FromError(err)
		if !ok{
                 panic("解析error错误")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}
	//fmt.Println(r.Message)
}
//////////////////////////////////////////////////////////////
service
/////////////////////////////////////////////////////////////////
package main
import(
	"context"
	"net"
	//"time"

	"all/grpc_test/proto"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/status"
	//"google.golang.org/grpc/codes"
)
type Server struct{
	//proto.UnimplementedGreeterServer
}
//ctx主要解决协程超时
func (s *Server) SayHello(ctx context.Context,in*proto.HelloRequest) (*proto.HelloReply, error) {
	//返回错误代码
	//return nil,status.Errorf(codes.NotFound,"记录未找到：%s",in.Name)
	//time.Sleep(5*time.Second)
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
////////////////////////////////////////////////////////////////////////
sql
////////////////////////////////////////////////////////////////////////
main.go
//////////////////////////////////////////////////////////////////////
package main

import(
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"database/sql"
)

type Product struct {
	gorm.Model//一些默认字段
	Code  sql.NullString//通过nullstring来设置零值问题
	Price uint
  }
func main()  {
	dsn := "sql路径"


	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  SlowThreshold:              time.Second,   // Slow SQL threshold
		  LogLevel:                   logger.Info, // Log level
		  IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
		  ParameterizedQueries:      true,           // Don't include params in the SQL log
		  Colorful:                  true,          // Disable color
		},)


		db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		  })
		  if err!=nil{
			panic(err)
		}

	  // 迁移 schema
  db.AutoMigrate(&Product{})//定义一个表结构将表结构直接生成对应的表 -migrations实例化一个空表此处一个有sql语句
  // Create
  //db.Create(&Product{Code:sql.NullString{"D42",true}, Price:100})//创建一个例子
  product := Product{
    Code:  sql.NullString{String: "D42", Valid: true},
    Price: 100,
}
db.Create(&product)
  // Read
  //var product Product
  db.First(&product, 1) // 根据整型主键查找
  db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
  // Update - 将 product 的 price 更新为 200
  db.Model(&product).Update("Price", 200)
  // Update - 更新多个字段
  db.Model(&product).Updates(Product{Price: 200, Code:sql.NullString{String: "D42", Valid: true}}) // 仅更新非零值字段 将code“”写入只是code没有改变将price设为0则price还是没有改变
  //如果我们去更新一个produce 只设置了price200 会导致其他值设为默认值
  //利用nullstring
  product= Product{
    Code:  sql.NullString{"",true},
    Price: 100,
}
  db.Model(&product).Updates(&product)
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
  // Delete - 删除 product
  db.Delete(&product, 1)
}
	  // Globally mode
	  
//设置全局的logger，这个logger在我们只想每个sql语句时候会打印一行sql
//sql是最重要的
