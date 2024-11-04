package main

import (
	"all/proto_bin/proto"
	"net/http"

	//"proto"

	"github.com/gin-gonic/gin"
)

func createGoods(c *gin.Context) {
	var mes struct {
		Name   string `json:"user"`
		Mess   string
		Number int
	}
	mes.Name = "bobby"
	mes.Mess = "测试mes"
	mes.Number = 3
	c.JSON(http.StatusOK, mes)
}
func createproto(c *gin.Context) {
	course := []string{"python", "go", "微服务"}
	user := proto.Teacher{
		Name:   "bobby",
		Course: course,
	}
	c.ProtoBuf(http.StatusOK, user)
}
func main() {
	// 实例化一个 gin 的 server 对象
	r := gin.Default() // 默认开启 logger 日志和 recovery 返回状态中间件
	// 定义分组路由
	good := r.Group("/goods")
	{
		good.GET("/morejson", createGoods)
		good.GET("/moreproto", createproto)

	}

	// 启动服务器
	r.Run(":8083")
}
