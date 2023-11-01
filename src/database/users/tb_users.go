package users

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	UserName  string `gorm:"Column:username;not null"`
	UserCode  string `gorm:"Column:usercode;UNIQUE;not null"`
	PassWord  string `gorm:"Column:password;not null"`
	UserImage string `gorm:"Column:userimge"`
	gorm.Model
}
