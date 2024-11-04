package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint64    `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
	CreatedBy string
	UpdatedBy string
	Version   uint32 `gorm:"default:1"`
}
