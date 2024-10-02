package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id          int64          `gorm:"primarykey;comment:用户ID"`
	UserCode    string         `gorm:"unique;not null;comment:用户编码"`
	UserNo      int64          `gorm:"unique;not null;comment:用户编号"`
	UserName    string         `gorm:"default:null;comment:用户名"`
	Nickname    string         `gorm:"default:null;comment:昵称"`
	Password    string         `gorm:"not null;comment:密码"`
	Salt        string         `gorm:"default:null;comment:盐值"`
	Email       string         `gorm:"default:null;comment:邮箱"`
	Phone       string         `gorm:"index;not null;comment:电话"`
	OpenId      string         `gorm:"default:null;comment:微信OpenID"`
	UnionId     string         `gorm:"default:null;comment:微信UnionID"`
	ClientIp    string         `gorm:"default:null;comment:客户端IP"`
	LastLoginAt time.Time      `gorm:"default:null;comment:最后登录时间"`
	CreatedBy   string         `gorm:"default:null;comment:创建者"`
	UpdatedBy   string         `gorm:"default:null;comment:更新者"`
	CreatedAt   time.Time      `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"comment:删除时间"`
	Version     int            `gorm:"default:1;comment:版本号"`
}

func (u *User) TableName() string {
	return "t_user"
}
