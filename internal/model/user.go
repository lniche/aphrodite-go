package model

import (
	"time"
)

type User struct {
	UserCode   string    `gorm:"unique;not null;comment:User code"`                                  // User code
	UserNo     uint64    `gorm:"unique;not null;comment:User number"`                                // User number
	Username   string    `gorm:"unique;default:null;comment:Username"`                               // Username
	Nickname   string    `gorm:"default:null;comment:Nickname"`                                      // Nickname
	Password   string    `gorm:"default:null;comment:Password"`                                      // Password
	Salt       string    `gorm:"default:null;comment:Salt"`                                          // Salt
	Email      string    `gorm:"default:null;comment:Email"`                                         // Email
	Phone      string    `gorm:"index;not null;comment:Phone"`                                       // Phone
	OpenId     string    `gorm:"default:null;comment:WeChat OpenID"`                                 // WeChat OpenID
	ClientIp   string    `gorm:"default:null;comment:Client IP"`                                     // Client IP
	LoginAt    time.Time `gorm:"default:null;comment:Login time"`                                    // Login time
	LoginToken string    `gorm:"default:null;comment:Login token"`                                   // Login token
	Avatar     string    `gorm:"default:null;comment:avatar"`                                        // Avatar
	Status     uint8    `gorm:"default:1;comment:0: Unactivated, 1: Active, 2: Frozen, 3: Deleted"` // User status
	BaseModel
}

func (u *User) TableName() string {
	return "t_user"
}
