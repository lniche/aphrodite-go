package model

import (
	"time"
)

type User struct {
	UserCode    string    `gorm:"unique;not null;comment:用户编码"`
	UserNo      int64     `gorm:"unique;not null;comment:用户编号"`
	UserName    string    `gorm:"default:null;comment:用户名"`
	Nickname    string    `gorm:"default:null;comment:昵称"`
	Password    string    `gorm:"not null;comment:密码"`
	Salt        string    `gorm:"default:null;comment:盐值"`
	Email       string    `gorm:"default:null;comment:邮箱"`
	Phone       string    `gorm:"index;not null;comment:电话"`
	OpenId      string    `gorm:"default:null;comment:微信OpenID"`
	UnionId     string    `gorm:"default:null;comment:微信UnionID"`
	ClientIp    string    `gorm:"default:null;comment:客户端IP"`
	LastLoginAt time.Time `gorm:"default:null;comment:最后登录时间"`
	BaseModel
}

func (u *User) TableName() string {
	return "t_user"
}
