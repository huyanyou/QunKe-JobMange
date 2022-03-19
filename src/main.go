package main

import (
	"Qunke_JobMange/src/client"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	// 创建管理员
	r.POST("/", client.CreateManger)
	// 删除管理员
	r.DELETE("/:positon", client.DeleteManger)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
