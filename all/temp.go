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
setnx 将获取和设置值变成原子性
有些业务逻辑进行查询操作时(特别是在根据某一字段DESC,取最大一笔).可以使用limit 1 或者 top 1 来终止
任何情况都不要用 select * from table ，用具体的字段列表替换"*"，不要返回用不到的字段,避免全盘扫描！
由于abc前面用了“%”，因此该查询必然走全表查询,除非必要(模糊查询需要包含abc)，否则不要在关键词前加%
通常使用 union all 或 union 的方式替换“or”会得到更好的效果。where子句中使用了or关键字,索引将被放弃使用。
在where子句中使用 IS NULL 或 IS NOT NULL 判断，索引将被放弃使用，会进行全表查询。
不要在where子句中的“=”左边进行函数、算数运算或其他表达式运算，否则系统将可能无法正确使用索引。
mysql查询只是用一个索引，因此如果where子句中已经使用了索引的话，那么order by中的列是不会使用索引。因此数据库默认排序可以符合要求情况下不要使用排序操作；
尽量不要包含多个列的排序，如果需要最好给这些列创建复合索引。
.Inner join 和 left join、right join、子查询代替in
exist & in 优化
in()适合B表比A表数据小的情况
exists()适合B表比A表数据大的情况
ttl -（获取时时间-开始时间）-时钟飘逸 =真正有效时间 
cap理论和base理论
mq发送->rpcketmq:发送half消息
ro->mo发送消息发送成功
mq->sql执行事务
mq->rocommit还是rockback
如果消息丢失
ro到mq查询
mq到本地事务查询
mq-》rock 返回事务状态
get_cat/indices 查看全部
///////////////////////////////////////////////////////////////////////////
put+id新建数据
put/account/_doc/1 //put是不允许不加id的
{
	数据
}
///////////////////////////////////////////////////////////////////////////
post不带id新建数据 没有就创建有就报错
post user/_doc/
{
	数据
}
get /account 查看index的基本信息
get/account/_source/1 只看_source的基本信息_source是嵌入在内的数据
Get account/_search   //查看所有信息通过request body查询
{
"query":{
"match_all":{}
       }
}
///////////////////////////////////////////////////////////////////////////
给已有的数据增加字段
post user/_update/1  不用update会被覆盖
{
"doc":{
	"age":18
    }
}
///////////////////////////////////////////////////////////////////////////
_mget批量获取
get user/_search   //form与size分页在数据量比较小的情况下可行
{
"query":{
	"match_all":{}
      },
	"from":4,
	"size":4,
}
///////////////////////////////////////////////////////////////////////////
get user/_search  //注意因为分词所以单词大小写不敏感
{
   “query”:{
	   "match":{
		   "address":"street"
		   }
	   }
}
///////////////////////////////////////////////////////////////////////////
倒排索引
match_phrase   //也会做分词 但是结果中要有所有分词，而且顺序要求一样
get user/_search  //注意因为分词所以单词大小写不敏感
{
   “query”:{
	   "match_phrase":{
		   "address":"street"
		   }
	   }
}
///////////////////////////////////////////////////////////////////////////
query_string 和match类似但是match要指定字段名 query_string不用他在所有字段查询
///////////////////////////////////////////////////////////////////////////
term查询这个不会分词查询
term 查询
用途：用于精确匹配字段值。term 查询不进行任何分词处理，直接查找与给定值完全匹配的文档。
适用场景：适用于过滤或搜索精确值，例如数字、日期、关键词等。
get/_search
{
	"query":{
	"term":{
	“user.id”{
	"value":"kimchy",
	"boost":1.0
		}
		}
		}
}
"boost":1.0 代表的是权重
///////////////////////////////////////////////////////////////////////////
range查询 -范围查询    //gt> //gte>=//it<//ite<=
Get/_search
{
"query":{
"range":{
"age":{	
"gte":10,
"it":20,
"boost":2.0
}
}
///////////////////////////////////////////////////////////////////////////
fuzzy模糊匹配，允许一定程度的拼写错误。
Get/_search
{ 
 "query":{
 "fuzzy":{
"user.id":{
  "value":"ki"
}
	 }
	}
}
///////////////////////////////////////////////////////////////////////////
bool 查询
must：所有指定的查询条件都必须匹配。相当于布尔逻辑中的 AND。
should：至少一个指定的查询条件必须匹配。相当于布尔逻辑中的 OR。
must_not：所有指定的查询条件都不能匹配。相当于布尔逻辑中的 NOT。
filter：所有指定的查询条件都必须匹配，但不影响相关性评分。适用于过滤操作。
get user/_search
{
  "query": {
    "bool": {
      "must": [
        { "range": { "price": { "gt": 100 } } },
        { "term": { "category": "Electronics" } }
      ]
    }
  }
}

	
