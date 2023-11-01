package server

import (
	"github.com/gin-gonic/gin"
	"myblog/src/database"
	"myblog/src/server/user"
)

func Sever() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//开启连接
	database.Dblaunch()
	defer database.Close()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	allGroup := r.Group("/api")
	{
		userGroup := allGroup.Group("/user")
		{
			userGroup.POST("/addUser", user.AddUser())
		}
	}
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":9797")
}
