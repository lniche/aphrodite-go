package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int64          `gorm:"primarykey"`
	UserCode  string         `gorm:"unique;not:null"`
	UserNo    int64          `gorm:"unique;not:null"`
	UserName  string         `gorm:"default:null"`
	Nickname  string         `gorm:"default:null"`
	Password  string         `gorm:"not null"`
	Salt      string         `gorm:"default null"`
	Email     string         `gorm:"default:null"`
	Phone     string         `gorm:"not:null"`
	OpenId    string         `gorm:"default:null"`
	UnionId   string         `gorm:"default:null"`
	CreatedBy string         `gorm:"default:null"`
	UpdatedBy string         `gorm:"default:null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Version   int            `gorm:"default:1"`
}

func (u *User) TableName() string {
	return "t_user"
}
