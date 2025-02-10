package models

import (
	"time"
	"gorm.io/gorm"
)

type Fight struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-" swaggerignore:"true"`
	EventID     uint           `gorm:"not null" json:"event_id"` // Foreign key to Event
	Fighter1    string         `gorm:"not null" json:"fighter_1"`
	Fighter2    string         `gorm:"not null" json:"fighter_2"`
	WeightClass string         `gorm:"not null" json:"weight_class"`
}