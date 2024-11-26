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
http
sql
docker
k8s
跨域问题
CORS的工作原理
跨域请求的场景
yuque.com/bobby-zpcyu/ggq3y6/ipym8u
/////////////////////////////////////////////////////////////
github.com/go-gorm/gorm
github.com/facebook/ent
github.com/facebook/sqlx

go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

https://gorm.io/zh_CN/docs/logger.html
///////////////////////////////////////////////
将proto反解回来
func init()先配置、
//////////////////////////////////////
zap
Info：记录信息级别的日志。
Debug：记录调试级别的日志。
Warn：记录警告级别的日志。
Error：记录错误级别的日志。
DPanic：记录严重错误级别的日志，并在开发环境中触发 panic。
Panic：记录错误级别的日志，并触发 panic。
Fatal：记录致命错误级别的日志，并终止程序。

商品  订单  用户 gorm sql一定要多学习

servec层 从sql里拿数据用gorm表单 然后写到proto定义的结构里 传给gin
同时写入sql也是用gorm 
同时前端到gin时用form json binding 表单限定
前到gin bool要设置成指针
返回的数据格式
mysql锁
redis锁
zook锁
grpc的长链接式编程 股票那种
type GoodsCategoryBrand struct {
	BaseModel
	ParentCategoryID int32      `gorm:"type:int;index_category_brand,unique;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"parent_category_id"`
	Category         Category   `gorm:"foreignKey:ParentCategoryID;references:ID" json:"-"`
	BrandsID         int32      `gorm:"type:int;index_category_brand,unique;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"brands_id"`
	Brands           Brands     `gorm:"foreignKey:BrandsID;references:ID" json:"-"`
}
result := global.DB.Preload("Brands").Where(&model.GoodsCategoryBrand{ParentCategoryID:req.Id}).Find(&categoryBrands)
type GoodsCategoryBrand struct{
	BaseModel
	ParentCategoryID int32      `gorm:"type:int;index_category_brand,unique"`
	Category         Category  

	BrandsID         int32      `gorm:"type:int;index_category_brand,unique"`
	Brands           Brands    
}
gorm里定义的数据格式什么时候用默认什么时候用数据库的

syntax = "proto3";

package example;

service StreamService {
  rpc SendMultipleMessages (stream MessageRequest) returns (MessageResponse) {}
  rpc ProcessMultipleMessages (stream SingleMessageRequest) returns (stream SingleMessageResponse) {}
}

message MessageRequest {
  string message = 1;
}

message MessageResponse {
  string result = 1;
}

message SingleMessageRequest {
  string message = 1;
  int32 id = 2;
}

message SingleMessageResponse {
  string processed_message = 1;
  bool success = 2;
}
客户端流：
服务端：需要一个 Recv 方法来接收客户端的多个请求，以及一个 SendAndClose 方法来发送最终的响应并关闭流。
客户端：需要一个 Send 方法来发送多个请求，以及一个 CloseAndRecv 方法来关闭发送方向的流并接收最终的响应。
服务端流：
服务端：需要一个 Send 方法来发送多个响应。
客户端：需要一个 Recv 方法来接收多个响应，以及一个 CloseSend 方法来关闭发送方向的流（虽然在服务端流中通常不需要调用 CloseSend）。
双向流：
服务端：需要一个 Send 方法来发送多个响应，以及一个 Recv 方法来接收多个请求。
客户端：需要一个 Send 方法来发送多个请求，以及一个 Recv 方法来接收多个响应。
总结
客户端流：客户端发送多个请求，服务器返回一个响应。
服务端流：客户端发送一个请求，服务器返回多个响应。
双向流：客户端和服务器都可以发送多个请求和响应。
