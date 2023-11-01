package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var DB *gorm.DB

func Dblaunch() (err error) {
	DB, err = gorm.Open("mssql", "sqlserver://sa:123456@localhost:1433?database=myblogdb")

	if err != nil {
		fmt.Println(err)
		fmt.Println("连接失败")
		return err
	}
	fmt.Println("连接成功")
	return DB.DB().Ping()
}

func Close() {
	DB.Close()

}
