package main

import (
	"net/http"

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
func main() {
	// 实例化一个 gin 的 server 对象
	r := gin.Default() // 默认开启 logger 日志和 recovery 返回状态中间件
	// 定义分组路由
	good := r.Group("/goods")
	{
		good.GET("/morejson", createGoods)
		good.POST("/add", createGoods)

	}

	// 启动服务器
	r.Run(":8083")
}
