package models

import (
	"time"
	"gorm.io/gorm"
)

type Event struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" swaggerignore:"true"`
	Name      string         `gorm:"not null" json:"name"`
	Date      string         `json:"date"`
	Location  string         `json:"location"`
	URL       string         `json:"url"`
}
