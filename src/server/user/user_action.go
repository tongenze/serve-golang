package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"myblog/src/database"
	"myblog/src/database/users"
)

func AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		p, _ := c.GetRawData()                 //获取到传过来的参数
		var u users.User                       //定义一个表类型的结构体存储参数
		_ = json.Unmarshal(p, &u)              //json反序列化进刚才的结构体 字段名会自动对应
		database.DB.AutoMigrate(&users.User{}) //自动进数据库建表 参数为定义的结构体地址
		res := database.DB.Create(&u)

		c.JSON(200, res)
	}
}
