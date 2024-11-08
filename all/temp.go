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
github.com/go-gorm/gorm
github.com/facebook/ent
github.com/facebook/sqlx

go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

https://gorm.io/zh_CN/docs/logger.html
///////////////////////////////////////////////
将proto反解回来
func init()先配置
