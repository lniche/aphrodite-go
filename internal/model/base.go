package model

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	CreatedBy string `gorm:"default:null;comment:创建者"`
	UpdatedBy string `gorm:"default:null;comment:更新者"`
	Version   int    `gorm:"default:1;comment:版本号"`
	Deleted   bool   `gorm:"column:is_deleted;default:false;comment:软删除标识，false表示未删除，true表示已删除"`
}
