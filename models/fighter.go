package models

import (
	"time"
	"gorm.io/gorm"
)

type Fighter struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `gorm:"unique;not null" json:"name"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Nickname  string         `json:"nickname"`
	Height    string         `json:"height"`
	Weight    string         `json:"weight"`
	Reach     string         `json:"reach"`
	Stance    string         `json:"stance"`
	Wins      int            `json:"wins"`
	Losses    int            `json:"losses"`
	Draws     int            `json:"draws"`
}