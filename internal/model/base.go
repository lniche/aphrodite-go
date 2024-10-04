package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	CreatedBy string
	UpdatedBy string
	Version   int  `gorm:"default:1"`
	Deleted   bool `gorm:"column:is_deleted;default:false"`
}
