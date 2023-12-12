package models

import (
	"gorm.io/gorm"
	"time"
)

type Link struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Alias     string
	Dest      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
