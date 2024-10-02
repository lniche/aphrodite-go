package model

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	CreatedBy string `gorm:"default:null;comment:创建者"`
	UpdatedBy string `gorm:"default:null;comment:更新者"`
	Version   int    `gorm:"default:1;comment:版本号"`
}
