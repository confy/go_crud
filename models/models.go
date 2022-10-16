package models

import (
	"time"

	"gorm.io/gorm"
)

type Data struct {
	gorm.Model
	Value     string `json:"value"`
	ID        uint   `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
