package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint           `gorm:"primarykey"`
	UserId    string         `gorm:"unique;not null"`
	UserCode  string         `gorm:"unique;not:null"`
	UserNo    uint           `gorm:"unique;not:null"`
	Username  string         `gorm:"default:null"`
	Nickname  string         `gorm:"default:null"`
	Password  string         `gorm:"not null"`
	Email     string         `gorm:"default:null"`
	Phone     string         `gorm:"default:null"`
	CreatedBy string         `gorm:"default:null"`
	UpdatedBy string         `gorm:"default:null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Version   int            `gorm:"default:0"`
}

func (u *User) TableName() string {
	return "t_user"
}
