package model

import (
	"time"
)

type User struct {
	UserCode   string    `gorm:"unique;not null;comment:用户编码"`
	UserNo     uint64    `gorm:"unique;not null;comment:用户编号"`
	Username   string    `gorm:"unique;default:null;comment:用户名"`
	Nickname   string    `gorm:"default:null;comment:昵称"`
	Password   string    `gorm:"not null;comment:密码"`
	Salt       string    `gorm:"default:null;comment:盐值"`
	Email      string    `gorm:"default:null;comment:邮箱"`
	Phone      string    `gorm:"index;not null;comment:电话"`
	OpenId     string    `gorm:"default:null;comment:微信OpenID"`
	ClientIp   string    `gorm:"default:null;comment:客户端IP"`
	LoginAt    time.Time `gorm:"default:null;comment:登录时间"`
	LoginToken string    `gorm:"default:null;comment:登录令牌"`
	BaseModel
}

func (u *User) TableName() string {
	return "t_user"
}
