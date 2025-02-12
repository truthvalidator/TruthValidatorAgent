package model

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	ID          string `gorm:"primaryKey"`
	Question    string `gorm:"type:text"`
	FirstAnswer string `gorm:"type:longtext"`
	Evidence    string `gorm:"type:longtext"`
	IsOk        bool   `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
